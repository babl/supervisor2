package main

import (
	"errors"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	"github.com/larskluge/babl-server/kafka"
	"github.com/larskluge/babl/bablmodule"
	bn "github.com/larskluge/babl/bablnaming"
	pb "github.com/larskluge/babl/protobuf"
	pbm "github.com/larskluge/babl/protobuf/messages"
	"golang.org/x/net/context"
)

type server struct {
	kafkaClient   *sarama.Client
	kafkaProducer *sarama.SyncProducer
}

type responses struct {
	channels map[uint64]chan *[]byte
	mux      sync.Mutex
}

const (
	Version                = "2.2.3"
	ModuleExecutionTimeout = 3 * time.Minute
	ModuleExecutionGrace   = 1 * time.Minute
	MaxGrpcMessageSize     = 1024 * 1024 * 1 // 1mb
)

var (
	debug    bool
	hostname string
	resp     responses

	Random = rand.New(rand.NewSource(time.Now().UnixNano()))
)

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
	resp = responses{channels: make(map[uint64]chan *[]byte)}

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

	data, err := s.request(ctx, in, async)

	if err != nil {
		return nil, err
	}
	if err := proto.Unmarshal(*data, out); err != nil {
		return nil, err
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

	rid := uint64(Random.Uint32())<<32 + uint64(Random.Uint32())

	key := hostname + "." + strconv.FormatUint(rid, 10)

	// Sends message to the babl module topic: e.g. "babl.larskluge.ImageResize.IO"
	topic := bn.RequestPathToTopic(MethodFromContext(ctx))
	module := bn.TopicToModule(topic)
	l := log.WithFields(log.Fields{"module": module, "rid": rid})

	l.WithFields(log.Fields{"message_size": len(msg)}).Debug("Send message to module")
	kafka.SendMessage(s.kafkaProducer, key, topic, &msg)

	if async {
		elapsed := float64(time.Since(start).Seconds() * 1000)
		l.WithFields(log.Fields{"duration_ms": elapsed}).Info("Request processed async")
		return &[]byte{}, nil
	}

	resp.mux.Lock()
	resp.channels[rid] = make(chan *[]byte, 1)
	resp.mux.Unlock()

	defer func() {
		resp.mux.Lock()
		delete(resp.channels, rid)
		resp.mux.Unlock()
	}()

	timeLeft := ModuleExecutionTimeout
	gracePeriodOver := false
	for {
		select {
		case data := <-resp.channels[rid]:
			elapsed := float64(time.Since(start).Seconds() * 1000)
			l.WithFields(log.Fields{"duration_ms": elapsed}).Info("Module responded")
			return data, nil
		case <-time.After(timeLeft):
			if gracePeriodOver {
				l.WithFields(log.Fields{"timeout": ModuleExecutionTimeout + ModuleExecutionGrace}).Error("Module did not respond in grace period either, timeout")
				return nil, errors.New("Module execution timed out")
			} else {
				l.WithFields(log.Fields{"timeout": ModuleExecutionTimeout}).Warn("Module did not respond in time, cancelling request execution")

				if err := s.BroadcastCancelRequest(module, rid); err != nil {
					l.WithError(err).Warn("Broadcasting cancel request failed")
				}

				// start over with grace period
				timeLeft = ModuleExecutionGrace
				gracePeriodOver = true
				continue
			}
		}
	}
}
