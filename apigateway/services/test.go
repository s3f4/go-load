package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/s3f4/go-load/apigateway/library/log"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
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

	instances, err := s.ir.GetFromTerraform()
	if err != nil {
		return err
	}

	var runTest models.RunTest
	runTest.TestID = test.ID
	runTest.StartTime = &startTime

	if err := s.rtr.Create(&runTest); err != nil {
		log.Errorf("TestService.Start: %v", err)
		return err
	}

	for index, instance := range instances {
		fmt.Println(instance)
		instanceCount := uint64(len(instances))
		requestPerInstance := test.RequestCount / instanceCount
		portion := index + 1
		event := models.Event{
			Event: models.REQUEST,
			Payload: models.RequestPayload{
				RunTestID:            runTest.ID,
				Portion:              fmt.Sprintf("%d/%d", portion, instanceCount),
				URL:                  test.URL,
				RequestCount:         requestPerInstance,
				Method:               test.Method,
				Payload:              test.Payload,
				GoroutineCount:       test.GoroutineCount,
				ExpectedResponseBody: test.ExpectedResponseBody,
				ExpectedResponseCode: test.ExpectedResponseCode,
				Headers:              test.Headers,
				TransportConfig:      test.TransportConfig,
			},
		}

		message, err := json.Marshal(event)
		if err != nil {
			return err
		}

		if err := s.queueService.Send("worker", message); err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}
