package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSupervisor2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Supervisor2 Suite")
}
