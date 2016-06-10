package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	pb "github.com/larskluge/babl/protobuf"
	pbm "github.com/larskluge/babl/protobuf/messages"
	"github.com/larskluge/babl/shared"
	"github.com/nneves/kafka-tools/bkconsumer"
	"github.com/nneves/kafka-tools/bkproducer"
	"golang.org/x/net/context"
)

type server struct{}

func main() {
	log.SetOutput(os.Stderr)
	log.SetFormatter(&log.JSONFormatter{})

	app := configureCli()
	app.Run(os.Args)
}

func run(listen string) {
	lis, err := net.Listen("tcp", listen)
	if err != nil {
		log.WithFields(log.Fields{"error": err, "listen": listen}).Fatal("Failed to listen at port")
	}

	modules := knownModules()

	s := NewServer()

	fmt.Println(modules)
	for _, module := range modules {
		m := shared.NewModule(module, false)
		pb.RegisterBinaryServer(m.GrpcServiceName(), s, &server{})
	}
	s.Serve(lis)
}

func (s *server) IO(ctx context.Context, in *pbm.BinRequest) (*pbm.BinReply, error) {
	start := time.Now()

	msg, err := proto.Marshal(in)
	if err != nil {
		return nil, err
	}

	randNbr := uint32(random(1, 999999))
	randStr := strconv.FormatUint(uint64(randNbr), 10)

	kafkaInboxProducer(randStr, []byte{})

	// Sends message to the babl module topic: e.g. "babl.larskluge.ImageResize.IO"
	topic := TopicFromMethod(MethodFromContext(ctx))
	kafkaTopicProducer(randStr, topic, msg)

	data := kafkaInboxConsumer(randStr)

	res := &pbm.BinReply{}

	if err := proto.Unmarshal(data, res); err != nil {
		return nil, err
	}

	elapsed := float64(time.Since(start).Seconds() * 1000)
	fmt.Printf("took %.3fs\n", elapsed)

	return res, nil
}

func (s *server) Ping(_ context.Context, in *pbm.Empty) (*pbm.Pong, error) {
	log.Info("ping")
	res := pbm.Pong{Val: "fake pong"}
	return &res, nil
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func kafkaInboxProducer(id string, value []byte) {
	topic := "inbox." + id
	fmt.Printf("Inbox -> ID=%q, Topic=%q, ValueSize=%q\r\n", id, topic, len(value))
	bkproducer.Producer(id, topic, value)
}

func kafkaTopicProducer(id, topic string, value []byte) {
	fmt.Printf("Topic -> ID=%q , Topic=%q, ValueSize=%q\r\n", id, topic, len(value))
	bkproducer.Producer(id, topic, value)
}

func kafkaInboxConsumer(id string) []byte {
	topic := "inbox." + id
	fmt.Printf("Inbox <- ID=%q Topic=%q\r\n", id, topic)
	_, value := bkconsumer.Consumer(topic)
	return value
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
