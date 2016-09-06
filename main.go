package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	"github.com/larskluge/babl-server/kafka"
	"github.com/larskluge/babl-storage/uploader"
	"github.com/larskluge/babl/bablmodule"
	pb "github.com/larskluge/babl/protobuf"
	pbm "github.com/larskluge/babl/protobuf/messages"
	"golang.org/x/net/context"
)

type server struct {
	kafkaClient   *sarama.Client
	kafkaProducer *sarama.SyncProducer
}

type responses struct {
	channels map[string]chan *[]byte
	mux      sync.Mutex
}

const (
	Version                    = "2.1.3"
	ModuleExecutionWaitTimeout = 5 * time.Minute
	UploadEndpoint             = "babl.sh:4443"

	MaxKafkaMessageSize = 1024 * 512        // 512kb
	MaxGrpcMessageSize  = 1024 * 1024 * 100 // 100mb
)

var debug bool
var hostname string
var resp responses

func main() {
	log.SetOutput(os.Stderr)
	log.SetFormatter(&log.JSONFormatter{})

	app := configureCli()
	app.Run(os.Args)
}

func run(listen, kafkaBrokers string, dbg bool) {
	debug = dbg
	if debug {
		log.SetLevel(log.DebugLevel)
	}

	hostname = Hostname()
	resp = responses{channels: make(map[string]chan *[]byte)}

	lis, err := net.Listen("tcp", listen)
	if err != nil {
		log.WithFields(log.Fields{"error": err, "listen": listen}).Fatal("Failed to listen at port")
	}

	s := server{}
	server := NewServer()

	brokers := strings.Split(kafkaBrokers, ",")
	clientID := "supervisor." + hostname
	s.kafkaClient = kafka.NewClient(brokers, clientID, debug)
	defer (*s.kafkaClient).Close()
	s.kafkaProducer = kafka.NewProducer(brokers, clientID+".producer")
	defer (*s.kafkaProducer).Close()

	newModulesChan := make(chan string)
	go discoverModules(s.kafkaClient, newModulesChan)
	go func() {
		for module := range newModulesChan {
			log.Infof("New Module Discovered: %s", module)
			m := bablmodule.New(module)
			pb.RegisterBinaryServer(m.GrpcServiceName(), server, &s)
		}
	}()
	go listenToModuleResponses(s.kafkaClient)

	log.Infof("Server started at %s", listen)
	server.Serve(lis)
}

func (s *server) IO(ctx context.Context, in *pbm.BinRequest) (*pbm.BinReply, error) {
	out := &pbm.BinReply{}
	_, async := in.Env["BABL_ASYNC"]

	if len(in.Stdin) > MaxKafkaMessageSize {
		upload, err := uploader.New(UploadEndpoint, bytes.NewReader(in.Stdin))
		check(err)
		log.WithFields(log.Fields{"blob_id": upload.Id, "blob_url": upload.Url}).Info("Store large payload externally")
		go func(upload *uploader.Upload) {
			success := upload.WaitForCompletion()
			if !success {
				log.WithFields(log.Fields{"blob_id": upload.Id, "blob_url": upload.Url}).Error("Large payload upload failed")
			}
		}(upload)
		in.Stdin = []byte{}
		in.PayloadUrl = upload.Url
	}

	data, err := s.request(ctx, in, async)

	if err != nil {
		return nil, err
	}
	if err := proto.Unmarshal(*data, out); err != nil {
		return nil, err
	}

	if !async {
		if out.PayloadUrl != "" {
			var err error
			out.Stdout, err = getPayload(out.PayloadUrl)
			check(err)
		}
	}

	return out, err
}

func (s *server) Ping(ctx context.Context, in *pbm.Empty) (*pbm.Pong, error) {
	out := &pbm.Pong{}
	data, err := s.request(ctx, in, false)
	if err != nil {
		return nil, err
	}
	if err := proto.Unmarshal(*data, out); err != nil {
		return nil, err
	}
	return out, err
}

func (s *server) request(ctx context.Context, in proto.Message, async bool) (*[]byte, error) {
	start := time.Now()

	msg, err := proto.Marshal(in)
	if err != nil {
		return nil, err
	}

	randNbr := uint32(random(1, 999999))
	rid := strconv.FormatUint(uint64(randNbr), 10)
	key := hostname + "." + rid

	// Sends message to the babl module topic: e.g. "babl.larskluge.ImageResize.IO"
	topic := TopicFromMethod(MethodFromContext(ctx))
	log.WithFields(log.Fields{"topic": topic, "key": key, "value size": len(msg)}).Debug("Send message to module")
	kafka.SendMessage(s.kafkaProducer, key, topic, &msg)

	if async {
		elapsed := float64(time.Since(start).Seconds() * 1000)
		log.WithFields(log.Fields{"duration_ms": elapsed, "rid": rid}).Info("Request processed async")
		return &[]byte{}, nil
	}

	resp.mux.Lock()
	resp.channels[rid] = make(chan *[]byte, 1)
	resp.mux.Unlock()

	select {
	case data := <-resp.channels[rid]:
		delete(resp.channels, rid)
		elapsed := float64(time.Since(start).Seconds() * 1000)
		log.WithFields(log.Fields{"duration_ms": elapsed, "rid": rid}).Info("Module responded")
		return data, nil
	case <-time.After(ModuleExecutionWaitTimeout):
		log.Error("Module execution timed out")
		return nil, errors.New("Module execution timed out")
	}
}

func getPayload(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
