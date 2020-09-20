package repository

import (
	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/gorm"
)

// ResponseRepository is used for processes on timescaledB
type ResponseRepository interface {
	DB() *gorm.DB
	List(query interface{}) ([]*models.Response, error)
}

type responseRepository struct {
	base BaseRepository
}

// NewResponseRepository returns new ResponseRepository instance
func NewResponseRepository() ResponseRepository {
	return &responseRepository{}
}

func (r *responseRepository) DB() *gorm.DB {
	return r.base.GetDB()
}

func (r *responseRepository) List(query interface{}) ([]*models.Response, error) {
	var responses []*models.Response

	if err := r.DB().Find(responses).Error; err != nil {
		return nil, err
	}

	return responses, nil
}
