package main

import (
	"os"
)

func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		panic("Please provide service name to start")
	}

	serviceName := args[1]

	if serviceName == "web" {
		startWebService()
		return

	} else if serviceName == "events" {
		startEventsService()
		return

	}
	panic("unimplemented")

}
