package main

import (
	"github.com/Shopify/sarama"
)

func knownModules() (mods []string) {
	modules := map[string]interface{}{}
	client, err := sarama.NewClient([]string{"localhost:9092"}, nil)
	check(err)
	defer client.Close()

	c, err := sarama.NewConsumerFromClient(client)
	check(err)
	defer c.Close()

	newestOffset, err := client.GetOffset("modules", 0, sarama.OffsetNewest)
	check(err)

	pc, err := c.ConsumePartition("modules", 0, sarama.OffsetOldest)
	check(err)
	defer c.Close()

	for message := range pc.Messages() {
		modules[string(message.Key)] = nil

		if message.Offset+1 >= newestOffset {
			break
		}
	}

	for key, _ := range modules {
		mods = append(mods, key)
	}
	return
}
