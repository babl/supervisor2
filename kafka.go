package main

import (
	"fmt"
	"strings"

	"github.com/nneves/kafka-tools/bkconsumer"
	"github.com/nneves/kafka-tools/bkproducer"
)

func TopicFromMethod(method string) string {
	return strings.Replace(method[1:], "/", ".", 1)
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
