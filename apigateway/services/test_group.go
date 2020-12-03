package services

import (
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
)

// TestGroupService creates tests
type TestGroupService interface {
	Start(*models.TestGroup) error
	Create(*models.TestGroup) error
	Get(*models.TestGroup) (*models.TestGroup, error)
	Update(*models.TestGroup) error
	Delete(*models.TestGroup) error
}

type testGroupService struct {
	ir           repository.InstanceRepository
	tgr          repository.TestGroupRepository
	rtr          repository.RunTestRepository
	queueService QueueService
	testService  TestService
}

// NewTestGroupService returns a testGroupService instance
func NewTestGroupService() TestGroupService {
	return &testGroupService{
		ir:           repository.NewInstanceRepository(),
		tgr:          repository.NewTestGroupRepository(),
		rtr:          repository.NewRunTestRepository(),
		queueService: NewRabbitMQService(),
		testService:  NewTestService(),
	}
}

func (s *testGroupService) Start(testGroup *models.TestGroup) error {
	for _, test := range testGroup.Tests {
		s.testService.Start(test)
	}
	return nil
}

// Create
func (s *testGroupService) Create(testGroup *models.TestGroup) error {
	return s.tgr.Create(testGroup)
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
