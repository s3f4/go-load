package main

import (
	"encoding/json"

	"github.com/s3f4/go-load/eventhandler/library/specs"
	"github.com/s3f4/go-load/eventhandler/models"
	"github.com/s3f4/go-load/eventhandler/repository"
	"github.com/s3f4/go-load/eventhandler/services"
	"github.com/s3f4/mu/log"
	"github.com/streadway/amqp"
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
	spec := &specs.ListenSpec{
		Consumer: "eventhandler",
		Queue:    "eventhandler",
		ProcessFunc: func(d *amqp.Delivery, exit chan<- struct{}) error {
			log.Infof("Received a message: %s", d.Body)
			var resp models.Response
			if err := json.Unmarshal(d.Body, &resp); err != nil {
				log.Error(err)
			}
			responseRepository.Create(&resp)
			return nil
		},
	}

	specs := []*specs.ListenSpec{
		spec,
	}

	service.Start(specs)
}
