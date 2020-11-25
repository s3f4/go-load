package repository

import (
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/gorm"
)

// TestRepository ..
type TestRepository interface {
	DB() *gorm.DB
	Create(*models.Test) error
	Update(*models.Test) error
	Delete(*models.Test) error
	Get(id uint) (*models.Test, error)
	List(*library.QueryBuilder, string, ...interface{}) ([]models.Test, int64, error)
}

type testRepository struct {
	base BaseRepository
}

var testRepositoryObject TestRepository

// NewTestRepository returns an testRepository object
func NewTestRepository() TestRepository {
	if testRepositoryObject == nil {
		testRepositoryObject = &testRepository{
			base: NewBaseRepository(MYSQL),
		}
	}
	return testRepositoryObject
}

func (r *testRepository) DB() *gorm.DB {
	return r.base.GetDB()
}

func (r *testRepository) Create(test *models.Test) error {
	return r.DB().Create(test).Error
}

func (r *testRepository) Update(test *models.Test) error {
	return r.DB().Session(&gorm.Session{FullSaveAssociations: true}).Updates(test).Error
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
	if err := r.DB().Preload("Headers").
		Preload("TestGroup").
		Preload("RunTests").
		Preload("TransportConfig").
		First(&testReq, id).Error; err != nil {
		return nil, err
	}
	return &testReq, nil
}

func (r *testRepository) List(query *library.QueryBuilder, conditionStr string, where ...interface{}) ([]models.Test, int64, error) {
	var testReq []models.Test
	var total int64
	if len(conditionStr) > 0 && where != nil {
		r.DB().Where(conditionStr, where...).Model(&testReq).Count(&total)
	} else {
		r.DB().Model(&testReq).Count(&total)
	}

	if err := query.SetDB(r.DB()).
		SetPreloads("Headers", "RunTests", "TransportConfig").
		SetModel(models.Test{}).
		SetWhere(conditionStr, where...).
		List(&testReq); err != nil {
		return nil, 0, err
	}

	return testReq, total, nil
}
