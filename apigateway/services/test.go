package services

import (
	"encoding/json"
	"fmt"

	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
	"github.com/s3f4/mu/log"
)

// TestService creates tests
type TestService interface {
	Start(models.RunConfig) error
}

type testService struct {
	ir           repository.InstanceRepository
	queueService QueueService
}

// NewTestService returns a testService instance
func NewTestService() TestService {
	return &testService{
		ir:           repository.NewInstanceRepository(),
		queueService: NewRabbitMQService(),
	}
}

func (s *testService) Start(runConfig models.RunConfig) error {
	iReq, err := s.ir.Get()
	log.Debug(fmt.Sprintf("%+v", iReq))

	if err != nil {
		return err
	}

	runConfig.InstanceCount = iReq.Configs[0].InstanceCount

	var requestPerInstance int
	if runConfig.InstanceCount != 0 {
		requestPerInstance = runConfig.RequestCount / runConfig.InstanceCount
	} else {
		requestPerInstance = runConfig.RequestCount
	}

	work := models.Work{
		Request:         requestPerInstance,
		URL:             runConfig.URL,
		GoroutineCount:  runConfig.GoroutineCount,
		TransportConfig: runConfig.TransportConfig,
	}

	message, err := json.Marshal(work)
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
