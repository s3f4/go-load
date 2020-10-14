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
	Get(id uint) (*models.RunTest, error)
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

func (r *runTestRepository) Insert(runTest *models.RunTest) error {
	return r.DB().Create(runTest).Error
}

func (r *runTestRepository) Delete(runTest *models.RunTest) error {
	return r.DB().Model(runTest).Delete(runTest).Error
}

func (r *runTestRepository) Get(id uint) (*models.RunTest, error) {
	var testReq models.RunTest
	if err := r.DB().First(&testReq, id).Error; err != nil {
		return nil, err
	}
	return &testReq, nil
}

func (r *runTestRepository) List() ([]models.RunTest, error) {
	var testReq []models.RunTest
	if err := r.DB().Find(&testReq).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return testReq, nil
}
