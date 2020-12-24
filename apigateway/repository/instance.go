package repository

import (
	"encoding/json"
	"errors"

	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/gorm"
)

// InstanceRepository ..
type InstanceRepository interface {
	Create(*models.InstanceConfig) error
	Delete(*models.InstanceConfig) error
	Get() (*models.InstanceConfig, error)
	GetFromTerraform() ([]models.InstanceTerraform, error)
}

type instanceRepository struct {
	db      *gorm.DB
	command library.Command
}

// NewInstanceRepository returns an instanceRepository object
func NewInstanceRepository(db *gorm.DB, command library.Command) InstanceRepository {
	return &instanceRepository{
		db:      db,
		command: command,
	}
}

func (r *instanceRepository) Create(instance *models.InstanceConfig) error {
	if err := r.db.Where("1=1").Delete(&models.InstanceConfig{}).Error; err != nil {
		return err
	}

	if err := r.db.Where("1=1").Delete(&models.Instance{}).Error; err != nil {
		return err
	}

	return r.db.Create(instance).Error
}

func (r *instanceRepository) Delete(instance *models.InstanceConfig) error {
	return r.db.Where("1=1").Delete(instance).Error
}

func (r *instanceRepository) Get() (*models.InstanceConfig, error) {
	var instanceReq models.InstanceConfig
	if err := r.db.Preload("Configs").Last(&instanceReq).Error; err != nil {
		return nil, err
	}
	return &instanceReq, nil
}

// GetFromTerraform
func (r *instanceRepository) GetFromTerraform() ([]models.InstanceTerraform, error) {
	output, err := r.command.Run("cd infra;terraform output -json workers")
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(output, &result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, errors.New("There is no instance")
	}

	var instances []models.InstanceTerraform
	for _, instance := range result {
		var instanceObj models.InstanceTerraform
		library.DecodeMap(instance, &instanceObj)
		instances = append(instances, instanceObj)
	}

	return instances, nil
}
