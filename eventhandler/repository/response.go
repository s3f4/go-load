package repository

import (
	"github.com/s3f4/go-load/eventhandler/models"
	"gorm.io/gorm"
)

// ResponseRepository is used for processes on timescaledB
type ResponseRepository interface {
	DB() *gorm.DB
	Insert(*models.Response) error
	Delete(*models.Response) error
	List(query interface{}) ([]*models.Response, error)
}

type responseRepository struct {
	base BaseRepository
}

// NewResponseRepository returns new ResponseRepository instance
func NewResponseRepository() ResponseRepository {
	return &responseRepository{
		base: NewBaseRepository(POSTGRES),
	}
}

func (r *responseRepository) DB() *gorm.DB {
	return r.base.GetDB()
}
func (r *responseRepository) Insert(response *models.Response) error {
	if err := r.DB().Create(response).Error; err != nil {
		return err
	}
	return nil
}

func (r *responseRepository) Delete(response *models.Response) error { return nil }

func (r *responseRepository) List(query interface{}) ([]*models.Response, error) {
	return nil, nil
}
