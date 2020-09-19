package repository

import (
	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/gorm"
)

// InstanceRepository ..
type InstanceRepository interface {
	Insert(*models.Instance) error
	Update(*models.Instance) error
	Delete(*models.Instance) error
	Get() (*models.Instance, error)
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

func (r *instanceRepository) Insert(instance *models.Instance) error {
	return r.DB().Create(instance).Error
}

func (r *instanceRepository) Update(instance *models.Instance) error {
	return r.DB().Model(instance).Update("test", "test").Error
}

func (r *instanceRepository) Delete(*models.Instance) error {
	return nil
}

func (r *instanceRepository) Get() (*models.Instance, error) {
	var instanceReq models.Instance
	if err := r.DB().Find(&instanceReq).Error; err != nil {
		return nil, err
	}
	return &instanceReq, nil
}
