package repository

import (
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/gorm"
)

// SettingsRepository ..
type SettingsRepository interface {
	Create(*models.Settings) error
	Delete(*models.Settings) error
	Update(*models.Settings) error
	Get(id uint) (*models.Settings, error)
	List(*library.QueryBuilder, string, ...interface{}) ([]models.Settings, int64, error)
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

func (r *setttingsRepository) Get(id uint) (*models.Settings, error) {
	var settings models.Settings
	if err := r.db.First(&settings, id).Error; err != nil {
		return nil, err
	}
	return &settings, nil
}

func (r *setttingsRepository) List(query *library.QueryBuilder, conditionStr string, where ...interface{}) ([]models.Settings, int64, error) {
	var settings []models.Settings
	var total int64
	if len(conditionStr) > 0 && where != nil {
		r.db.Where(conditionStr, where...).Model(&settings).Count(&total)
	} else {
		r.db.Model(&settings).Count(&total)
	}

	if err := query.SetDB(r.db).
		SetModel(models.Settings{}).
		SetWhere(conditionStr, where...).
		List(&settings); err != nil {
		return nil, 0, err
	}

	return settings, total, nil
}
