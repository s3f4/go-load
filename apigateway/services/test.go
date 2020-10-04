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
	Start(models.TestConfig) error
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

func (s *testService) Start(testConfig models.TestConfig) error {
	instanceConfig, err := s.ir.Get()
	log.Debug(fmt.Sprintf("%+v", instanceConfig))

	if err != nil {
		return err
	}

	for _, test := range testConfig.Tests {
		for _, instance := range instanceConfig.Configs {
			requestPerInstance := test.RequestCount / instance.InstanceCount

			event := models.Event{
				Event: models.REQUEST,
				Payload: models.RequestPayload{
					URL:                  test.URL,
					RequestCount:         requestPerInstance,
					Method:               test.Method,
					Payload:              test.Payload,
					GoroutineCount:       test.GoroutineCount,
					ExpectedResponseBody: test.ExpectedResponseBody,
					ExpectedResponseCode: test.ExpectedResponseCode,
					TransportConfig:      test.TransportConfig,
				},
			}

			message, err := json.Marshal(event)
			if err != nil {
				return err
			}

			log.Info(string(message))

			for i := 0; i < instance.InstanceCount; i++ {
				if err := s.queueService.Send("worker", message); err != nil {
					return err
				}
			}
		}

	}

	return nil
}
