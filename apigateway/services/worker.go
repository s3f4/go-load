package services

import (
	"encoding/json"

	"github.com/s3f4/go-load/apigateway/models"
)

// WorkerService service makes request and runs
// Worker container properly
type WorkerService interface {
	Run(models.RunConfig) error
}

type workerService struct {
	queueService QueueService
}

// NewWorkerService returns new worker service object
func NewWorkerService() WorkerService {
	return &workerService{
		queueService: NewRabbitMQService("amqp://user:password@queue:5672/"),
	}
}

func (s *workerService) Run(runConfig models.RunConfig) error {
	requestPerInstance := runConfig.RequestCount / runConfig.InstanceCount
	queueObj := map[string]interface{}{
		"request": requestPerInstance,
		"url":     runConfig.URL,
	}

	message, err := json.Marshal(queueObj)
	if err != nil {
		return err
	}

	for i := 0; i < runConfig.InstanceCount; i++ {
		if err := s.queueService.Send("queue", message); err != nil {
			return err
		}
	}

	return nil
}
