package main

import (
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/larskluge/babl-server-kafka/kafka"
	"github.com/larskluge/babl-server-kafka/kafka/oldconsumer"
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

func kafkaInboxProducer(id string, value []byte) {
	topic := "inbox." + id
	fmt.Printf("Inbox -> ID=%q, Topic=%q, ValueSize=%q\r\n", id, topic, len(value))
	bablkafka.Producer(id, topic, value)
}

func kafkaTopicProducer(id, topic string, value []byte) {
	fmt.Printf("Topic -> ID=%q , Topic=%q, ValueSize=%q\r\n", id, topic, len(value))
	bablkafka.Producer(id, topic, value)
}

func kafkaInboxConsumer(id string) []byte {
	topic := "inbox." + id
	fmt.Printf("Inbox <- ID=%q Topic=%q\r\n", id, topic)
	_, value := oldconsumer.Consumer(topic)
	return value
}
