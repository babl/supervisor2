package main

import (
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Hostname() string {
	h, err := os.Hostname()
	check(err)
	return h
}
