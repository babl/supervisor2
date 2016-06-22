package main

import (
	"strings"

	"github.com/Shopify/sarama"
	log "github.com/Sirupsen/logrus"
	"github.com/larskluge/babl-server/kafka"
)

// NewKafkaClient returns new Kafka Client connection
func NewKafkaClient(brokerStr string) *sarama.Client {
	brokers := strings.Split(brokerStr, ",")
	client, err := sarama.NewClient(brokers, nil)
	check(err)
	return &client
}

// TopicFromMethod replace topic char '/' with '.'
func TopicFromMethod(method string) string {
	return strings.Replace(method[1:], "/", ".", 1)
}

func kafkaTopicProducer(id, topic string, value []byte) {
	log.Debugf("Topic -> ID=%q , Topic=%q, ValueSize=%q", id, topic, len(value))
	kafka.Producer(id, topic, value, kafka.ProducerOptions{Verbose: debug})
}

func kafkaInboxConsumer(id string) []byte {
	topic := "out." + id
	log.Debugf("Consumer: Payload Out <- ID=%q Topic=%q", id, topic)
	_, value := kafka.Consumer(topic, kafka.ConsumerOptions{Verbose: debug})
	return value
}
