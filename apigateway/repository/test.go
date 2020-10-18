package repository

import (
	"fmt"

	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/gorm"
)

// TestRepository ..
type TestRepository interface {
	DB() *gorm.DB
	Insert(*models.Test) error
	Update(*models.Test) error
	Delete(*models.Test) error
	Get(id uint) (*models.Test, error)
	List() ([]models.Test, error)
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

func (r *testRepository) Insert(test *models.Test) error {
	return r.DB().Create(test).Error
}

func (r *testRepository) Update(test *models.Test) error {
	return r.DB().Model(test).Updates(test).Error
}

func (r *testRepository) Delete(test *models.Test) error {
	err := r.DB().Where("test_id=?", test.ID).Delete(&models.TransportConfig{}).Error
	if err != nil {
		return err
	}
	return r.DB().Where("id=?", test.ID).Delete(test).Error
}

func (r *testRepository) Get(id uint) (*models.Test, error) {
	var testReq models.Test
	if err := r.DB().Take(&testReq).Error; err != nil {
		return nil, err
	}
	return &testReq, nil
}

func (r *testRepository) List() ([]models.Test, error) {
	var testReq []models.Test
	if err := r.DB().Preload("Headers").Preload("RunTests").Preload("TransportConfig").Find(&testReq).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return testReq, nil
}
