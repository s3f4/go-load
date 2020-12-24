package main

import (
	"github.com/s3f4/go-load/eventhandler/models"
	"github.com/s3f4/go-load/eventhandler/repository"
	"github.com/s3f4/go-load/eventhandler/services"
	"gorm.io/gorm"
)

func main() {

	var postgresConn *gorm.DB
	postgresConn = repository.Connect(repository.POSTGRES)
	postgresConn.AutoMigrate(&models.Response{})

	queues := []string{"eventhandler"}
	s := services.NewRabbitMQService(postgresConn)
	s.QueueDeclare(queues[0])
	s.QueueDeclare("worker")

	service := services.NewListener(postgresConn)
	service.Start(queues...)
}
