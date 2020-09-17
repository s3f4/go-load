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

func (r *instanceRepository) Insert(instance *models.Instance) error {
	return r.base.Insert(instance)
}

func (r *instanceRepository) Update(*models.Instance) error {
	return nil
}

func (r *instanceRepository) Delete(*models.Instance) error {
	return nil
}

func (r *instanceRepository) Get(id int) error {
	return nil
}
