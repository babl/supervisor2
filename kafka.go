package main

import (
	"strings"
)

// TopicFromMethod replace topic char '/' with '.'
func TopicFromMethod(method string) string {
	return strings.Replace(method[1:], "/", ".", 1)
}
