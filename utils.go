package main

import (
	"math/rand"
	"os"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func Hostname() string {
	h, err := os.Hostname()
	check(err)
	return h
}

func init() {
	rand.Seed(time.Now().Unix())
}
