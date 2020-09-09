package repository

import "github.com/s3f4/go-load/apigateway/models"

// InstanceRepository ..
type InstanceRepository interface {
	Insert(*models.Instance) error
	Update(*models.Instance) error
	Delete(*models.Instance) error
	Get(id int) error
}

type instanceRepository struct {
	base BaseRepository
}

// NewInstanceRepository returns an instanceRepository object
func NewInstanceRepository() InstanceRepository {
	return &instanceRepository{
		base: NewBaseRepository(),
	}
}

func (ir *instanceRepository) Insert(instance *models.Instance) error {
	return ir.base.Insert(instance)
}
func (ir *instanceRepository) Update(*models.Instance) error { return nil }
func (ir *instanceRepository) Delete(*models.Instance) error { return nil }
func (ir *instanceRepository) Get(id int) error              { return nil }
