package services

import (
	"encoding/json"
	"fmt"

	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
	"github.com/s3f4/mu/log"
)

// WorkerService service makes request and runs
// Worker container properly
type WorkerService interface {
	Run(models.RunConfig) error
}

type workerService struct {
	ir           repository.InstanceRepository
	queueService QueueService
}

// NewWorkerService returns new worker service object
func NewWorkerService() WorkerService {

	return &workerService{
		queueService: NewRabbitMQService(),
		ir:           repository.NewInstanceRepository(),
	}
}

func (s *workerService) Run(runConfig models.RunConfig) error {
	iReq, err := s.ir.Get()
	log.Info(fmt.Sprintf("%+v", iReq))

	if err != nil {
		return err
	}

	runConfig.InstanceCount = iReq.InstanceCount

	var requestPerInstance int
	if iReq.InstanceCount != 0 {
		requestPerInstance = runConfig.RequestCount / iReq.InstanceCount
	} else {
		requestPerInstance = runConfig.RequestCount
	}

	queueObj := map[string]interface{}{
		"request": requestPerInstance,
		"url":     runConfig.URL,
	}

	message, err := json.Marshal(queueObj)
	if err != nil {
		return err
	}
	log.Info(string(message))

	for i := 0; i < runConfig.InstanceCount; i++ {
		if err := s.queueService.Send("worker", message); err != nil {
			return err
		}
	}

	return nil
}
