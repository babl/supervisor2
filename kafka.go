package main

import (
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/larskluge/babl-server/kafka"
	"github.com/larskluge/babl-server/kafka/singleconsumer"
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
	if debug {
		fmt.Printf("Inbox -> ID=%q, Topic=%q, ValueSize=%q\r\n", id, topic, len(value))
	}
	kafka.Producer(id, topic, value, kafka.ProducerOptions{Verbose: debug})
}

func kafkaTopicProducer(id, topic string, value []byte) {
	if debug {
		fmt.Printf("Topic -> ID=%q , Topic=%q, ValueSize=%q\r\n", id, topic, len(value))
	}
	kafka.Producer(id, topic, value, kafka.ProducerOptions{Verbose: debug})
}

func kafkaInboxConsumer(id string) []byte {
	topic := "inbox." + id
	if debug {
		fmt.Printf("Inbox <- ID=%q Topic=%q\r\n", id, topic)
	}
	_, value := singleconsumer.Consumer(topic, debug)
	return value
}
