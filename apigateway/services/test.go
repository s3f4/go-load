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

	instanceCount := uint64(len(instances))

	if test.RequestCount < instanceCount {
		for i := uint64(0); i < test.RequestCount; i++ {
			event := setEvent(&runTest, test, 1, test.RequestCount, i+1)
			message, err := json.Marshal(event)
			if err != nil {
				log.Error(err)
				return err
			}

			if err := s.queueService.Send("worker", message); err != nil {
				log.Error(err)
				return err
			}
		}
	} else {
		for i := range instances {
			requestPerInstance := test.RequestCount / instanceCount
			event := setEvent(&runTest, test, requestPerInstance, instanceCount, uint64(i+1))

			// add remain RequestCount to RequestCount of  last event
			if len(instances) == i+1 {
				event.Payload.(*models.RequestPayload).RequestCount = event.Payload.(*models.RequestPayload).RequestCount + uint64((test.RequestCount - (requestPerInstance * instanceCount)))
			}

			message, err := json.Marshal(event)
			if err != nil {
				log.Error(err)
				return err
			}

			if err := s.queueService.Send("worker", message); err != nil {
				log.Error(err)
				return err
			}
		}
	}

	return nil
}

func setEvent(runTest *models.RunTest, test *models.Test, requestPerInstance, instanceOrRequestCount, portion uint64) *models.Event {
	return &models.Event{
		Event: models.REQUEST,
		Payload: models.RequestPayload{
			RunTestID:            runTest.ID,
			Portion:              fmt.Sprintf("%d/%d", portion, instanceOrRequestCount),
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
}
