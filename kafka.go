package main

import (
	"io/ioutil"
	stdlog "log"
	"os"
	"strings"

	"github.com/Shopify/sarama"
	log "github.com/Sirupsen/logrus"
	// "github.com/larskluge/babl-server/kafka"
	// "github.com/larskluge/babl-server/kafka/singleconsumer"
)

// NewKafkaClient returns new Kafka Client connection
func NewKafkaClient(brokerStr string) *sarama.Client {
	logger := stdlog.New(os.Stderr, "", stdlog.LstdFlags)
	if debug {
		logger.SetOutput(ioutil.Discard)
	} else {
		logger.SetOutput(os.Stderr)
	}
	sarama.Logger = logger

	brokers := strings.Split(brokerStr, ",")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.ClientID = "supervisor." + hostname
	// config.Producer.Partitioner = sarama.NewManualPartitioner
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	client, err := sarama.NewClient(brokers, config)
	check(err)
	return &client
}

func NewKafkaProducer(client *sarama.Client) *sarama.SyncProducer {
	producer, err := sarama.NewSyncProducerFromClient(*client)
	check(err)
	return &producer
}

func NewKafkaConsumer(client sarama.Client) *sarama.Consumer {
	consumer, err := sarama.NewConsumerFromClient(client)
	check(err)
	return &consumer
}

// TopicFromMethod replace topic char '/' with '.'
func TopicFromMethod(method string) string {
	return strings.Replace(method[1:], "/", ".", 1)
}

func sendMessage(producer *sarama.SyncProducer, key, topic string, value []byte) {
	log.Debugf("Topic -> Key=%q , Topic=%q, ValueSize=%q", key, topic, len(value))

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(value),
	}
	if key != "" {
		msg.Key = sarama.StringEncoder(key)
	}
	// defer func() {
	// 	if err := producer.Close(); err != nil {
	// 		log.Panicf("Producer: Failed to close Kafka producer cleanly: %+q", err)
	// 	}
	// }()
	_, _, err := (*producer).SendMessage(msg)
	check(err)
}

func consumeOutTopic(consumer *sarama.Consumer, key string) []byte {
	topic := "out." + key
	log.Debugf("Consumer: Payload Out <- Key=%q Topic=%q", key, topic)
	// _, value := singleconsumer.Consumer(topic, debug)

	pc, err := (*consumer).ConsumePartition(topic, 0, sarama.OffsetOldest)
	check(err)
	defer pc.AsyncClose()
	log.WithFields(log.Fields{"partition": 0, "topic": topic, "key": key}).Info("Consume Partition")
	for msg := range pc.Messages() {
		log.Info("new message received")
		return msg.Value
	}
	return []byte{}
}
