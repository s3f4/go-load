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
	Insert(*models.Test) error
	Get(*models.Test) (*models.Test, error)
	Update(*models.Test) error
	Delete(*models.Test) error
	List() ([]models.Test, error)
}

type testService struct {
	ir           repository.InstanceRepository
	tr           repository.TestRepository
	queueService QueueService
}

// NewTestService returns a testService instance
func NewTestService() TestService {
	return &testService{
		ir:           repository.NewInstanceRepository(),
		tr:           repository.NewTestRepository(),
		queueService: NewRabbitMQService(),
	}
}

// Insert
func (s *testService) Insert(test *models.Test) error {
	return s.tr.Insert(test)
}

// Get
func (s *testService) Get(test *models.Test) (*models.Test, error) {
	return s.tr.Get(test.ID)
}

// Update
func (s *testService) Update(test *models.Test) error {
	return s.tr.Update(test)
}

// Delete
func (s *testService) Delete(test *models.Test) error {
	return s.tr.Delete(test)
}

// List
func (s *testService) List() ([]models.Test, error) {
	return s.tr.List()
}

func (s *testService) Run(testID uint) error {
	instanceConfig, err := s.ir.Get()
	if err != nil {
		return err
	}

	test, err := s.tr.Get(testID)

	if err != nil {
		return err
	}

	for _, instance := range instanceConfig.Configs {
		requestPerInstance := test.RequestCount / uint64(instance.Count)

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
