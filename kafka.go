package main

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/larskluge/babl-server/kafka"
)

// TopicFromMethod replace topic char '/' with '.'
func TopicFromMethod(method string) string {
	return strings.Replace(method[1:], "/", ".", 1)
}

func kafkaInboxConsumer(id string) []byte {
	topic := "out." + id
	log.Debugf("Consumer: Payload Out <- ID=%q Topic=%q", id, topic)
	_, value := kafka.Consumer(topic, kafka.ConsumerOptions{Verbose: debug})
	return value
}
