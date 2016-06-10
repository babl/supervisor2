package main

import (
	"github.com/Shopify/sarama"
)

func discoverModules(client sarama.Client, ch chan string) {
	c, err := sarama.NewConsumerFromClient(client)
	check(err)
	defer c.Close()

	pc, err := c.ConsumePartition("modules", 0, sarama.OffsetOldest)
	check(err)
	defer pc.Close()

	modules := map[string]interface{}{}
	for message := range pc.Messages() {
		module := string(message.Key)

		_, known := modules[module]

		if !known {
			modules[module] = nil
			ch <- module
		}
	}
}
