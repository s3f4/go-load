package services

import (
	"encoding/json"
	"fmt"

	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
	"github.com/s3f4/mu/log"
)

// TestGroupService creates tests
type TestGroupService interface {
	Start(*models.TestGroup) error
	Insert(*models.TestGroup) error
	Get(*models.TestGroup) (*models.TestGroup, error)
	Update(*models.TestGroup) error
	Delete(*models.TestGroup) error
	List() ([]models.TestGroup, error)
}

type testGroupService struct {
	ir           repository.InstanceRepository
	tgr          repository.TestGroupRepository
	queueService QueueService
}

// NewTestGroupService returns a testGroupService instance
func NewTestGroupService() TestGroupService {
	return &testGroupService{
		ir:           repository.NewInstanceRepository(),
		tgr:          repository.NewTestGroupRepository(),
		queueService: NewRabbitMQService(),
	}
}

func (s *testGroupService) Start(testGroup *models.TestGroup) error {
	instanceConfig, err := s.ir.Get()
	log.Debug(fmt.Sprintf("%+v", instanceConfig))

	if err != nil {
		return err
	}

	for _, test := range testGroup.Tests {
		for _, instance := range instanceConfig.Configs {
			fmt.Println(instance)
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
					fmt.Println(err)
					return err
				}
			}
		}

	}

	return nil
}

// Insert
func (s *testGroupService) Insert(testGroup *models.TestGroup) error {
	return s.tgr.Insert(testGroup)
}

// Get
func (s *testGroupService) Get(testGroup *models.TestGroup) (*models.TestGroup, error) {
	return s.tgr.Get(testGroup.ID)
}

// Update
func (s *testGroupService) Update(testGroup *models.TestGroup) error {
	return s.tgr.Update(testGroup)
}

// Delete
func (s *testGroupService) Delete(testGroup *models.TestGroup) error {
	return s.tgr.Delete(testGroup)
}

// List
func (s *testGroupService) List() ([]models.TestGroup, error) {
	return s.tgr.List()
}
