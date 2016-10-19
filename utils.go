package main

import (
	"os"
	"strings"
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

func TopicFromRequestPath(method string) string {
	return strings.Replace(method[1:], "/", ".", 1)
}
