package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Kafka", func() {
	Context("#TopicFromMethod", func() {
		It("converts a simple request string to a Kafka topic", func() {
			Expect(TopicFromMethod("/babl.larskluge.Telegram/IO")).To(Equal("babl.larskluge.Telegram.IO"))
		})
	})
})
