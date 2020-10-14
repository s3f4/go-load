package services

import (
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
)

// RunTestService creates runTests
type RunTestService interface {
	Insert(*models.RunTest) error
	Get(*models.RunTest) (*models.RunTest, error)
	Delete(*models.RunTest) error
	List() ([]models.RunTest, error)
}

type runTestService struct {
	rtr repository.RunTestRepository
}

// NewRunTestService returns a runTestService instance
func NewRunTestService() RunTestService {
	return &runTestService{
		rtr: repository.NewRunTestRepository(),
	}
}

// Insert
func (s *runTestService) Insert(runTest *models.RunTest) error {
	return s.rtr.Insert(runTest)
}

// Get
func (s *runTestService) Get(runTest *models.RunTest) (*models.RunTest, error) {
	return s.rtr.Get(runTest.ID)
}

// Delete
func (s *runTestService) Delete(runTest *models.RunTest) error {
	return s.rtr.Delete(runTest)
}

// List
func (s *runTestService) List() ([]models.RunTest, error) {
	return s.rtr.List()
}
