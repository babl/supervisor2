package main

import (
	"strings"
)

func TopicFromMethod(method string) string {
	return strings.Replace(method[1:], "/", ".", 1)
}
