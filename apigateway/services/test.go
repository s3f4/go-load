package services

import (
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
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
