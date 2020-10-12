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
	Start(*models.TestConfig) error
	Insert(*models.TestConfig) error
	Get(*models.TestConfig) (*models.TestConfig, error)
	Update(*models.TestConfig) error
	Delete(*models.TestConfig) error
	UpdateTest(*models.Test) error
	DeleteTest(*models.Test) error
	List() ([]models.TestConfig, error)
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

func (s *testService) Start(testConfig *models.TestConfig) error {
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

// Insert
func (s *testService) Insert(config *models.TestConfig) error {
	return s.tr.Insert(config)
}

// Get
func (s *testService) Get(config *models.TestConfig) (*models.TestConfig, error) {
	return s.tr.Get()
}

// Update
func (s *testService) Update(config *models.TestConfig) error {
	return s.tr.Update(config)
}

// Delete
func (s *testService) Delete(config *models.TestConfig) error {
	return s.tr.Delete(config)
}

// DeleteTest
func (s *testService) DeleteTest(test *models.Test) error {
	return s.tr.DeleteTest(test)
}

// UpdateTest
func (s *testService) UpdateTest(test *models.Test) error {
	return s.tr.UpdateTest(test)
}

// List
func (s *testService) List() ([]models.TestConfig, error) {
	return s.tr.List()
}
