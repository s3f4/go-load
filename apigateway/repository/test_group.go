package repository

import (
	"fmt"

	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/gorm"
)

// TestGroupRepository ..
type TestGroupRepository interface {
	DB() *gorm.DB
	Insert(*models.TestGroup) error
	Update(*models.TestGroup) error
	Delete(*models.TestGroup) error
	Get() (*models.TestGroup, error)
	List() ([]models.TestGroup, error)
}

type testGroupRepository struct {
	base BaseRepository
}

// NewTestGroupRepository returns an testGroupRepository object
func NewTestGroupRepository() TestGroupRepository {
	return &testGroupRepository{
		base: NewBaseRepository(MYSQL),
	}
}

func (r *testGroupRepository) DB() *gorm.DB {
	return r.base.GetDB()
}

func (r *testGroupRepository) Insert(testGroup *models.TestGroup) error {
	return r.DB().Create(testGroup).Error
}

func (r *testGroupRepository) Update(testGroup *models.TestGroup) error {
	return r.DB().Model(testGroup).Updates(testGroup).Error
}

func (r *testGroupRepository) Delete(testGroup *models.TestGroup) error {
	return r.DB().Model(testGroup).Delete(testGroup).Error
}

func (r *testGroupRepository) Get() (*models.TestGroup, error) {
	var testReq models.TestGroup
	if err := r.DB().Take(&testReq).Error; err != nil {
		return nil, err
	}
	return &testReq, nil
}

func (r *testGroupRepository) List() ([]models.TestGroup, error) {
	var testReq []models.TestGroup
	if err := r.DB().Preload("Tests").Find(&testReq).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return testReq, nil
}
