//go:generate go-bindata data/...

package main

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	pbm "github.com/larskluge/babl/protobuf/messages"
	"github.com/larskluge/babl/shared"
	"github.com/nneves/kafka-tools/bkconsumer"
	"github.com/nneves/kafka-tools/bkproducer"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct{}

func main() {
	log.SetOutput(os.Stderr)
	log.SetFormatter(&log.JSONFormatter{})

	app := configureCli()
	app.Run(os.Args)
}

func run(listen string) {
	modules := knownModules()

	lis, err := net.Listen("tcp", listen)
	if err != nil {
		log.WithFields(log.Fields{"error": err, "listen": listen}).Fatal("Failed to listen at port")
	}

	certPEMBlock, _ := Asset("data/server.pem")
	keyPEMBlock, _ := Asset("data/server.key")
	cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	check(err)
	creds := credentials.NewServerTLSFromCert(&cert)
	opts := []grpc.ServerOption{grpc.Creds(creds)}

	s := grpc.NewServer(opts...)

	fmt.Println(modules)
	for _, module := range modules {
		m := shared.NewModule(module, false)
		RegisterBinaryServer(m.GrpcServiceName(), s, &server{})
	}
	s.Serve(lis)
}

func (s *server) IO(_ context.Context, method string, in *pbm.BinRequest) (*pbm.BinReply, error) {
	start := time.Now()

	msg, err := proto.Marshal(in)
	check(err)

	randNbr := uint32(random(1, 999999))
	randStr := strconv.FormatUint(uint64(randNbr), 10)

	topic := strings.Replace(method[1:], "/", ".", 1)
	fmt.Printf("Request URL: %q -> topic: %s\n", method, topic)

	kafkaInboxProducer(randStr, []byte{})

	// Sends message to the babl module topic: "babl.larskluge.ImageResize.IO"
	kafkaTopicProducer(randStr, topic, msg)

	data := kafkaInboxConsumer(randStr)

	// fmt.Fprintf(w, "Received Response: %q\r\n", datastr)

	res := &pbm.BinReply{}

	if err := proto.Unmarshal(data, res); err != nil {
		panic(err)
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
