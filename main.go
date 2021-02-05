package main

import (
	"github.com/YJ-dev/go-server/rest"
	"github.com/YJ-dev/go-server/socket"
)

func main() {
	runType("socket")
}

func runType(connType string) {
	if connType == "rest" {
		rest.Run()
	} else {
		socket.Run()
	}
}
