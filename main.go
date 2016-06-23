package main

import (
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	"github.com/larskluge/babl-server/kafka"
	pb "github.com/larskluge/babl/protobuf"
	pbm "github.com/larskluge/babl/protobuf/messages"
	"github.com/larskluge/babl/shared"
	"golang.org/x/net/context"
)

type server struct {
	kafkaClient  sarama.Client
	knownModules map[string]bool
}

var debug bool
var hostname string

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

	lis, err := net.Listen("tcp", listen)
	if err != nil {
		log.WithFields(log.Fields{"error": err, "listen": listen}).Fatal("Failed to listen at port")
	}

	s := server{}
	server := NewServer()

	brokers := strings.Split(kafkaBrokers, ",")
	s.kafkaClient = kafka.NewClient(brokers, debug)
	defer s.kafkaClient.Close()

	newModulesChan := make(chan string)
	go discoverModules(s.kafkaClient, newModulesChan)
	go func() {
		for module := range newModulesChan {
			log.Infof("New Module Discovered: %s", module)
			m := shared.NewModule(module, false)
			pb.RegisterBinaryServer(m.GrpcServiceName(), server, &s)
		}
	}()

	log.Infof("Server started at %s", listen)
	server.Serve(lis)
}

func (s *server) IO(ctx context.Context, in *pbm.BinRequest) (*pbm.BinReply, error) {
	out := &pbm.BinReply{}
	data, err := s.request(ctx, in)
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
	data, err := s.request(ctx, in)
	if err != nil {
		return nil, err
	}
	if err := proto.Unmarshal(*data, out); err != nil {
		return nil, err
	}
	return out, err
}

func (s *server) request(ctx context.Context, in proto.Message) (*[]byte, error) {
	start := time.Now()

	msg, err := proto.Marshal(in)
	if err != nil {
		return nil, err
	}

	randNbr := uint32(random(1, 999999))
	randStr := strconv.FormatUint(uint64(randNbr), 10)
	key := hostname + "." + randStr

	// Sends message to the babl module topic: e.g. "babl.larskluge.ImageResize.IO"
	topic := TopicFromMethod(MethodFromContext(ctx))
	producer := kafka.NewProducer(s.kafkaClient)
	defer producer.Close()
	log.Debugf("Topic -> Key=%q , Topic=%q, ValueSize=%q", key, topic, len(msg))
	kafka.SendMessage(producer, key, topic, msg)

	data := kafkaInboxConsumer(key)

	elapsed := float64(time.Since(start).Seconds() * 1000)
	log.Infof("took %.3fs\n", elapsed)

	return &data, nil
}
