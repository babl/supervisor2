package main

import (
	"github.com/Shopify/sarama"
	log "github.com/Sirupsen/logrus"
	"github.com/larskluge/babl-server/kafka"
	u "github.com/larskluge/babl-server/utils"
)

func listenToModuleResponses(client *sarama.Client) {
	topic := "supervisor." + hostname
	log.Debug("Consuming from supervisor topic")
	ch := make(chan *kafka.ConsumerData)
	go kafka.Consume(client, topic, ch)
	for msg := range ch {
		log.WithFields(log.Fields{"key": msg.Key}).Debug("Response received from module exec")

		rid, err := u.ParseRid(msg.Key)
		check(err)

		resp.mux.Lock()
		channel, ok := resp.channels[rid]
		resp.mux.Unlock()
		if ok {
			channel <- &msg.Value
			close(channel)
		}
		msg.Processed <- "success"
	}
	panic("listenToModuleResponses: Lost connection to Kafka")
}
