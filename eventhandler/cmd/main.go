package main

import (
	"github.com/s3f4/go-load/eventhandler/models"
	"github.com/s3f4/go-load/eventhandler/repository"
	"github.com/s3f4/go-load/eventhandler/services"
)

func main() {
	models := []interface{}{
		&models.Response{},
	}
	baseRepo := repository.NewBaseRepository(repository.POSTGRES)
	baseRepo.Migrate(models...)

	queues := []string{"eventhandler"}
	s := services.NewRabbitMQService()
	s.QueueDeclare(queues[0])
	s.QueueDeclare("worker")

	service := services.NewListener()
	service.Start(queues...)
}
