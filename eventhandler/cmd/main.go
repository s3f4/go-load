package main

import (
	"github.com/s3f4/go-load/eventhandler"
	"github.com/s3f4/mu/log"
)

func main() {
	queues := []string{"eventhandler", "worker"}

	s := eventhandler.NewRabbitMQService()
	for _, queue := range queues {
		log.Info(queue)
		s.QueueDeclare(queue)
	}

	service := eventhandler.NewListener()
	service.Start(queues...)
}
