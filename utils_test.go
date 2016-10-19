package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils", func() {
	Context("#TopicFromMethod", func() {
		It("converts a request path to a Kafka topic", func() {
			Expect(TopicFromRequestPath("/babl.larskluge.Telegram/IO")).To(Equal("babl.larskluge.Telegram.IO"))
		})
	})
})
