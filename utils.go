package main

import (
	"os"
	"regexp"
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

func RequestPathToTopic(method string) string {
	return strings.Replace(method[1:], "/", ".", 1)
}

func TopicToModuleName(topic string) string {
	parts := strings.Split(topic, ".")
	owner := parts[1]
	camel := parts[2]

	name := regexp.MustCompile(`[A-Z]+`).ReplaceAllStringFunc(camel, func(m string) string {
		return "-" + strings.ToLower(m)
	})
	name = strings.TrimLeft(name, "-")

	return owner + "/" + name
}

func ModuleNameToTopic(module string, meta bool) string {
	function := "IO"
	if meta {
		function = "meta"
	}

	parts := strings.Split(module, "/")
	owner := parts[0]
	name := strings.Replace(strings.Title(parts[1]), "-", "", -1)

	return "babl." + owner + "." + name + "." + function
}
