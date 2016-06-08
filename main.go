//go:generate go-bindata data/...

package main

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/golang/protobuf/proto"
	pb "github.com/larskluge/babl/protobuf"
	pbm "github.com/larskluge/babl/protobuf/messages"
	"github.com/larskluge/babl/shared"
	"github.com/nneves/kafka-tools/bkconsumer"
	"github.com/nneves/kafka-tools/bkproducer"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct{}

var command string

func main() {
	log.SetOutput(os.Stderr)
	log.SetFormatter(&log.JSONFormatter{})

	app := configureCli()
	app.Run(os.Args)
}

func configureCli() (app *cli.App) {
	app = cli.NewApp()
	app.Usage = "Server for a Babl Module"
	app.Version = "0.3.0"
	app.Action = defaultAction
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "module, m",
			Usage:  "Module to serve",
			EnvVar: "BABL_MODULE",
		},
		cli.StringFlag{
			Name:   "cmd",
			Usage:  "Command to be executed",
			Value:  "cat",
			EnvVar: "BABL_COMMAND",
		},
		cli.IntFlag{
			Name:   "port",
			Usage:  "Port for server to be started on",
			EnvVar: "PORT",
			Value:  4444,
		},
	}
	return
}

func defaultAction(c *cli.Context) {
	module := c.String("module")
	if module == "" {
		cli.ShowAppHelp(c)
		os.Exit(1)
	} else {
		if !shared.CheckModuleName(module) {
			log.WithFields(log.Fields{"module": module}).Fatal("Module name format incorrect")
		}
		command = c.String("cmd")
		address := fmt.Sprintf(":%d", c.Int("port"))

		log.Warn("Start module")

		lis, err := net.Listen("tcp", address)
		if err != nil {
			log.WithFields(log.Fields{"error": err, "address": address}).Fatal("Failed to listen at port")
		}

		certPEMBlock, _ := Asset("data/server.pem")
		keyPEMBlock, _ := Asset("data/server.key")
		cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
		check(err)
		creds := credentials.NewServerTLSFromCert(&cert)
		opts := []grpc.ServerOption{grpc.Creds(creds)}

		s := grpc.NewServer(opts...)
		m := shared.NewModule(module, false)
		pb.RegisterBinaryServer(m.GrpcServiceName(), s, &server{})
		s.Serve(lis)
	}
}

func (s *server) IO(ctx context.Context, in *pbm.BinRequest) (*pbm.BinReply, error) {
	start := time.Now()

	msg, err := proto.Marshal(in)
	check(err)

	randNbr := uint32(random(1, 999999))
	randStr := strconv.FormatUint(uint64(randNbr), 10)

	modulePath := "larskluge.image-resize"
	fmt.Printf("Request URL: %q\r\n", modulePath)

	kafkaInboxProducer(randStr, []byte{})

	// Sends message to the babl module topic: "larskluge.image-resize"
	kafkaTopicProducer(randStr, modulePath, msg)

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

func (s *server) Ping(ctx context.Context, in *pbm.Empty) (*pbm.Pong, error) {
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
