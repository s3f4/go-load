package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
	"github.com/s3f4/mu/log"
)

// TestService creates tests
type TestService interface {
	Start(test *models.Test) error
}

type testService struct {
	ir           repository.InstanceRepository
	tr           repository.TestRepository
	rtr          repository.RunTestRepository
	queueService QueueService
}

// NewTestService returns a testService instance
func NewTestService() TestService {
	return &testService{
		ir:           repository.NewInstanceRepository(),
		tr:           repository.NewTestRepository(),
		rtr:          repository.NewRunTestRepository(),
		queueService: NewRabbitMQService(),
	}
}

func (s *testService) Start(test *models.Test) error {
	startTime := time.Now()

	instanceConfig, err := s.ir.Get()
	if err != nil {
		return err
	}

	var runTest models.RunTest
	runTest.TestID = test.ID
	runTest.StartTime = &startTime

	if err := s.rtr.Create(&runTest); err != nil {
		return err
	}

	for _, instance := range instanceConfig.Configs {
		requestPerInstance := test.RequestCount / uint64(instance.Count)

		event := models.Event{
			Event: models.REQUEST,
			Payload: models.RequestPayload{
				RunTestID:            runTest.ID,
				URL:                  test.URL,
				RequestCount:         requestPerInstance,
				Method:               test.Method,
				Payload:              test.Payload,
				GoroutineCount:       test.GoroutineCount,
				ExpectedResponseBody: test.ExpectedResponseBody,
				ExpectedResponseCode: test.ExpectedResponseCode,
				TransportConfig:      test.TransportConfig,
				Headers:              test.Headers,
			},
		}

		message, err := json.Marshal(event)
		if err != nil {
			return err
		}

		log.Info(string(message))

		for i := 0; i < instance.Count; i++ {
			if err := s.queueService.Send("worker", message); err != nil {
				fmt.Println(err)
				return err
			}
		}

	}

	return nil
}
