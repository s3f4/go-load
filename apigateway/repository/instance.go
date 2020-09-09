package repository

import "github.com/jinzhu/gorm"

// InstanceRepository ..
type InstanceRepository interface {
	Insert()
	Update()
	Delete()
	Get()
}

type instanceRepository struct {
	BaseRepository
}

// NewInstanceRepository returns an instanceRepository object
func NewInstanceRepository(db *gorm.DB) InstanceRepository {
	return &instanceRepository{}
}

func (ir *instanceRepository) Insert() {}
func (ir *instanceRepository) Update() {}
func (ir *instanceRepository) Delete() {}
func (ir *instanceRepository) Get()    {}
