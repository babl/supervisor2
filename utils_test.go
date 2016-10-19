package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils", func() {
	Context("#RequestPathToTopic", func() {
		It("converts a request path to a Kafka topic", func() {
			Expect(RequestPathToTopic("/babl.larskluge.Telegram/IO")).To(Equal("babl.larskluge.Telegram.IO"))
		})
	})
})
