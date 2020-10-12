package repository

import (
	"fmt"

	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/gorm"
)

// TestRepository ..
type TestRepository interface {
	DB() *gorm.DB
	Insert(*models.TestConfig) error
	Update(*models.TestConfig) error
	Delete(*models.TestConfig) error
	UpdateTest(*models.Test) error
	DeleteTest(*models.Test) error
	Get() (*models.TestConfig, error)
	List() ([]models.TestConfig, error)
}

type testRepository struct {
	base BaseRepository
}

// NewTestRepository returns an testRepository object
func NewTestRepository() TestRepository {
	return &testRepository{
		base: NewBaseRepository(MYSQL),
	}
}

func (r *testRepository) DB() *gorm.DB {
	return r.base.GetDB()
}

func (r *testRepository) Insert(testConfig *models.TestConfig) error {
	return r.DB().Create(testConfig).Error
}

func (r *testRepository) Update(testConfig *models.TestConfig) error {
	return r.DB().Where("1=1").Delete(testConfig).Error
}

func (r *testRepository) Delete(testConfig *models.TestConfig) error {
	return r.DB().Where("1=1").Delete(testConfig).Error
}

func (r *testRepository) UpdateTest(test *models.Test) error {
	return r.DB().Model(test).Updates(test).Error
}

func (r *testRepository) DeleteTest(test *models.Test) error {
	err := r.DB().Where("test_id=?", test.ID).Delete(&models.TransportConfig{}).Error
	if err != nil {
		return err
	}
	return r.DB().Where("id=?", test.ID).Delete(test).Error
}

func (r *testRepository) Get() (*models.TestConfig, error) {
	var testReq models.TestConfig
	if err := r.DB().Last(&testReq).Error; err != nil {
		return nil, err
	}
	return &testReq, nil
}

func (r *testRepository) List() ([]models.TestConfig, error) {
	var testReq []models.TestConfig
	if err := r.DB().Preload("Tests").Find(&testReq).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return testReq, nil
}
