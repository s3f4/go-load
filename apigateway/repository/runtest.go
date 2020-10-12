package repository

import (
	"fmt"

	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/gorm"
)

// RunTestRepository ..
type RunTestRepository interface {
	DB() *gorm.DB
	Insert(*models.RunTest) error
	Delete(*models.RunTest) error
	Get() (*models.RunTest, error)
	List() ([]models.RunTest, error)
}

type runTestRepository struct {
	base BaseRepository
}

// NewRunTestRepository returns an testRepository object
func NewRunTestRepository() RunTestRepository {
	return &runTestRepository{
		base: NewBaseRepository(MYSQL),
	}
}

func (r *runTestRepository) DB() *gorm.DB {
	return r.base.GetDB()
}

func (r *runTestRepository) Insert(testGroup *models.RunTest) error {
	return r.DB().Create(testGroup).Error
}

func (r *runTestRepository) Delete(testGroup *models.RunTest) error {
	return r.DB().Model(testGroup).Delete(testGroup).Error
}

func (r *runTestRepository) Get() (*models.RunTest, error) {
	var testReq models.RunTest
	if err := r.DB().Take(&testReq).Error; err != nil {
		return nil, err
	}
	return &testReq, nil
}

func (r *runTestRepository) List() ([]models.RunTest, error) {
	var testReq []models.RunTest
	if err := r.DB().Preload("Tests").Find(&testReq).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return testReq, nil
}
