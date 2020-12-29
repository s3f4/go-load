package repository

import (
	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/gorm"
)

// SettingsRepository ..
type SettingsRepository interface {
	Create(*models.Settings) error
	Delete(*models.Settings) error
	Update(*models.Settings) error
	Get(key models.SettingsKey) (*models.Settings, error)
}

type setttingsRepository struct {
	db *gorm.DB
}

// NewSettingsRepository returns an testRepository object
func NewSettingsRepository(db *gorm.DB) SettingsRepository {
	return &setttingsRepository{
		db: db,
	}
}

func (r *setttingsRepository) Create(settings *models.Settings) error {
	return r.db.Create(settings).Error
}

func (r *setttingsRepository) Delete(settings *models.Settings) error {
	return r.db.Model(settings).Delete(settings).Error
}

func (r *setttingsRepository) Update(settings *models.Settings) error {
	return r.db.Model(settings).Updates(settings).Error
}

func (r *setttingsRepository) Get(key models.SettingsKey) (*models.Settings, error) {
	var settings models.Settings
	if err := r.db.Where("setting = ?", key).First(&settings).Error; err != nil {
		return nil, err
	}
	return &settings, nil
}
