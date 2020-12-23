package repository

import (
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/gorm"
)

// ResponseRepository is used for processes on timescaledB
type ResponseRepository interface {
	DB() *gorm.DB
	List(*library.QueryBuilder, string, ...interface{}) ([]models.Response, int64, error)
}

type responseRepository struct {
	db *gorm.DB
}

// NewResponseRepository returns new ResponseRepository instance
func NewResponseRepository(db *gorm.DB) ResponseRepository {
	return &responseRepository{
		db: db,
	}
}

func (r *responseRepository) DB() *gorm.DB {
	return r.db
}

func (r *responseRepository) List(query *library.QueryBuilder, conditionStr string, where ...interface{}) ([]models.Response, int64, error) {

	var responses []models.Response
	var total int64
	if len(conditionStr) > 0 && where != nil {
		r.DB().Where(conditionStr, where...).Model(&responses).Count(&total)
	} else {
		r.DB().Model(&responses).Count(&total)
	}

	if err := query.SetDB(r.DB()).
		SetModel(models.Response{}).
		SetWhere(conditionStr, where...).
		List(&responses); err != nil {
		return nil, 0, err
	}

	return responses, total, nil
}
