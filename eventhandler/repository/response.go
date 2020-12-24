package repository

import (
	"github.com/s3f4/go-load/eventhandler/models"
	"gorm.io/gorm"
)

// ResponseRepository is used for processes on timescaledB
type ResponseRepository interface {
	Create(*models.Response) error
}

type responseRepository struct {
	db *gorm.DB
}

// NewResponseRepository returns new ResponseRepository instance
func NewResponseRepository(conn *gorm.DB) ResponseRepository {
	return &responseRepository{
		db: conn,
	}
}

func (r *responseRepository) Create(response *models.Response) error {
	return r.db.Create(response).Error
}
