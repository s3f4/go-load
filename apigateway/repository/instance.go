package repository

import (
	"encoding/json"
	"errors"

	"github.com/mitchellh/mapstructure"
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/library/log"
	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/gorm"
)

// InstanceRepository ..
type InstanceRepository interface {
	DB() *gorm.DB
	Create(*models.InstanceConfig) error
	Delete(*models.InstanceConfig) error
	Get() (*models.InstanceConfig, error)
	GetFromTerraform() ([]models.InstanceTerraform, error)
}

type instanceRepository struct {
	base    BaseRepository
	command library.Command
}

var instanceRepositoryObject InstanceRepository

// NewInstanceRepository returns an instanceRepository object
func NewInstanceRepository(command library.Command) InstanceRepository {
	if instanceRepositoryObject == nil {
		instanceRepositoryObject = &instanceRepository{
			base:    NewBaseRepository(MYSQL),
			command: command,
		}
	}
	return instanceRepositoryObject
}

func (r *instanceRepository) DB() *gorm.DB {
	return r.base.GetDB()
}

func (r *instanceRepository) Create(instance *models.InstanceConfig) error {
	if err := r.DB().Where("1=1").Delete(&models.InstanceConfig{}).Error; err != nil {
		return err
	}

	if err := r.DB().Where("1=1").Delete(&models.Instance{}).Error; err != nil {
		return err
	}

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
		decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			Metadata: nil,
			Result:   &instanceObj,
			TagName:  "json",
		})
		if err != nil {
			log.Errorf("mapstructrure.decode", err)
			return nil, err
		}

		if err := decoder.Decode(instance); err != nil {
			log.Errorf("worker.start", err)
			return nil, err
		}

		instances = append(instances, instanceObj)
	}

	return instances, nil
}
