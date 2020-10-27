package repository

import (
	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/gorm"
)

// ResponseRepository is used for processes on timescaledB
type ResponseRepository interface {
	DB() *gorm.DB
	List(runTestID uint) ([]*models.Response, error)
}

type responseRepository struct {
	base BaseRepository
}

var responseRepositoryObject ResponseRepository

// NewResponseRepository returns new ResponseRepository instance
func NewResponseRepository() ResponseRepository {
	if responseRepositoryObject == nil {
		responseRepositoryObject = &responseRepository{
			base: NewBaseRepository(POSTGRES),
		}
	}
	return responseRepositoryObject
}

func (r *responseRepository) DB() *gorm.DB {
	return r.base.GetDB()
}

func (r *responseRepository) List(runTestID uint) ([]*models.Response, error) {
	var responses []*models.Response
	if err := r.DB().Where("run_test_id", runTestID).Find(&responses).Error; err != nil {
		return nil, err
	}

	return responses, nil
}
