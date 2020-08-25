package main

import (
	"github.com/s3f4/go-load/eventhandler"
)

func main() {
	uri := "amqp://user:password@queue:5672/"
	queues := []string{"worker"}
	service := eventhandler.NewListener(uri)
	service.Start(queues...)
}
