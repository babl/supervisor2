package main

import (
	"github.com/Shopify/sarama"
	log "github.com/Sirupsen/logrus"
	"github.com/larskluge/babl-server/kafka"
)

var responses = make(map[string]chan []byte)

func listenToModuleResponses(client sarama.Client) {
	topic := "supervisor." + hostname
	for {
		log.Info("Consuming from supervisor topic")
		key, value := kafka.Consumer(topic, kafka.ConsumerOptions{Verbose: debug})
		channel, ok := responses[key]
		if ok {
			channel <- value
		}
	}
}
