package main

import (
	"github.com/s3f4/go-load/eventhandler"
)

func main() {
	queues := []string{"eventhandler"}
	s := eventhandler.NewRabbitMQService()
	s.QueueDeclare(queues[0])
	s.QueueDeclare("worker")

	service := eventhandler.NewListener()
	service.Start(queues...)
}
