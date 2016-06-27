package main

import (
	"github.com/Shopify/sarama"
	log "github.com/Sirupsen/logrus"
	"github.com/larskluge/babl-server/kafka"
)

var responses = make(map[string]chan []byte)

func listenToModuleResponses(client *sarama.Client) {
	topic := "supervisor." + hostname
	log.Debug("Consuming from supervisor topic")
	ch := make(chan kafka.ConsumerData)
	go kafka.Consume(client, topic, ch)
	for msg := range ch {
		log.WithFields(log.Fields{"key": msg.Key}).Debug("Response received from module exec")
		channel, ok := responses[msg.Key]
		if ok {
			channel <- msg.Value
			close(channel)
		}
	}
}
