package main

import (
	"github.com/s3f4/go-load/eventhandler"
)

func main() {
	queues := []string{"eventhandler"}
	service := eventhandler.NewListener()
	service.Start(queues...)
}
