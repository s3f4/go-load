package main

import (
	"github.com/s3f4/go-load/eventhandler/app"
	"github.com/s3f4/go-load/eventhandler/models"
	"github.com/s3f4/go-load/eventhandler/repository"
	"github.com/s3f4/go-load/eventhandler/services"
	"gorm.io/gorm"
)

func main() {

	var postgresConn *gorm.DB
	postgresConn = repository.Connect(repository.POSTGRES)
	postgresConn.AutoMigrate(&models.Response{})
	responseRepository := repository.NewResponseRepository(postgresConn)

	queues := []string{"eventhandler"}
	s := services.NewQueueService()
	s.QueueDeclare(queues[0])
	s.QueueDeclare("worker")

	service := services.NewListener()
	specs := app.GetSpecs(responseRepository)

	service.Start(specs)
}
