package main

import "github.com/s3f4/go-load/eventhandler/services"

func main() {
	queues := []string{"eventhandler"}
	s := services.NewRabbitMQService()
	s.QueueDeclare(queues[0])
	s.QueueDeclare("worker")

	service := services.NewListener()
	service.Start(queues...)
}
