package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/library/log"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
	"github.com/streadway/amqp"
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
	instances, err := s.ir.GetFromTerraform()
	if err != nil {
		return err
	}

	startTime := time.Now()
	var runTest models.RunTest
	runTest.TestID = test.ID
	runTest.StartTime = &startTime

	if err := s.rtr.Create(&runTest); err != nil {
		log.Errorf("TestService.Start: %v", err)
		return err
	}
	runTest.Test = test
	instanceCount := uint64(len(instances))

	if test.RequestCount < instanceCount {
		for i := uint64(0); i < test.RequestCount; i++ {
			event := setEvent(&runTest, 1, test.RequestCount, i+1)
			if err := s.sendMessage(event); err != nil {
				return err
			}
		}
	} else {
		for i := range instances {
			requestPerInstance := test.RequestCount / instanceCount
			// add remain RequestCount to RequestCount of  last event
			if len(instances) == i+1 {
				requestPerInstance = requestPerInstance + uint64((test.RequestCount - (requestPerInstance * instanceCount)))
			}

			event := setEvent(&runTest, requestPerInstance, instanceCount, uint64(i+1))

			if err := s.sendMessage(event); err != nil {
				return err
			}
		}
	}

	return s.waitQueue(&runTest, test.RequestCount, instanceCount)
}

func (s *testService) waitQueue(runTest *models.RunTest, requestCount, instanceCount uint64) error {
	// Declare Queue for this runTest
	queue := fmt.Sprintf("collect_%d_%d", runTest.Test.ID, runTest.ID)
	if err := s.queueService.Declare(queue); err != nil {
		return err
	}

	var payloads []*models.CollectPayload

	s.queueService.Listen(&ListenSpec{
		Queue:    queue,
		Consumer: fmt.Sprintf("apigateway_test_%d", runTest.Test.ID),
		ProcessFunc: func(d *amqp.Delivery, exit chan<- struct{}) error {
			var event models.Event
			if err := json.Unmarshal(d.Body, &event); err != nil {
				return err
			}

			var payload models.CollectPayload
			if err := library.DecodeMap(event.Payload, &payload); err != nil {
				return err
			}

			payloads = append(payloads, &payload)
			if s.checkPayloads(runTest, requestCount, instanceCount, payloads) {
				exit <- struct{}{}
			}
			return nil
		},
	})

	if err := s.queueService.Delete(queue); err != nil {
		return err
	}

	return nil
}

func (s *testService) checkPayloads(
	runTest *models.RunTest,
	requestCount, instanceCount uint64,
	payloads []*models.CollectPayload,
) bool {
	var total uint64
	if requestCount < instanceCount {
		total = requestCount
	} else {
		total = instanceCount
	}

	if int(total) != len(payloads) {
		return false
	}

	for _, payload := range payloads {
		if payload.RunTest.Passed == false {
			runTest.Passed = false
		}
		currentEndTime := runTest.EndTime
		if payload.RunTest.EndTime.After(*currentEndTime) {
			runTest.EndTime = payload.RunTest.EndTime
		}
	}

	if err := s.rtr.Update(runTest); err != nil {
		log.Errorf("RunTest Update error: %v\n", err)
	}

	return true
}

func (s *testService) sendMessage(event *models.Event) error {
	message, err := json.Marshal(event)
	if err != nil {
		log.Error(err)
		return err
	}

	if err := s.queueService.Send("worker", message); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func setEvent(runTest *models.RunTest, requestPerInstance, instanceOrRequestCount, portion uint64) *models.Event {
	return &models.Event{
		Event: models.REQUEST,
		Payload: models.RequestPayload{
			Portion:      fmt.Sprintf("%d/%d", portion, instanceOrRequestCount),
			RequestCount: requestPerInstance,
			RunTest:      runTest,
			Test:         runTest.Test,
		},
	}
}
