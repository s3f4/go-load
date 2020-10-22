package repository

import (
	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/gorm"
)

// InstanceRepository ..
type InstanceRepository interface {
	DB() *gorm.DB
	Create(*models.InstanceConfig) error
	Delete(*models.InstanceConfig) error
	Get() (*models.InstanceConfig, error)
}

type instanceRepository struct {
	base BaseRepository
}

// NewInstanceRepository returns an instanceRepository object
func NewInstanceRepository() InstanceRepository {
	return &instanceRepository{
		base: NewBaseRepository(MYSQL),
	}
}

func (r *instanceRepository) DB() *gorm.DB {
	return r.base.GetDB()
}

func (r *instanceRepository) Create(instance *models.InstanceConfig) error {
	return r.DB().Create(instance).Error
}

func (r *instanceRepository) Delete(instance *models.InstanceConfig) error {
	return r.DB().Where("1=1").Delete(instance).Error
}

func (r *instanceRepository) Get() (*models.InstanceConfig, error) {
	var instanceReq models.InstanceConfig
	if err := r.DB().Preload("Configs").Last(&instanceReq).Error; err != nil {
		return nil, err
	}
	return &instanceReq, nil
}
