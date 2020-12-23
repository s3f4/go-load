package repository

import (
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// TestRepository ..
type TestRepository interface {
	Create(*models.Test) error
	Update(*models.Test) error
	Delete(*models.Test) error
	Get(id uint) (*models.Test, error)
	List(*library.QueryBuilder, string, ...interface{}) ([]models.Test, int64, error)
}

type testRepository struct {
	db *gorm.DB
}

var testRepositoryObject TestRepository

// NewTestRepository returns an testRepository object
func NewTestRepository(db *gorm.DB) TestRepository {
	if testRepositoryObject == nil {
		testRepositoryObject = &testRepository{
			db: db,
		}
	}
	return testRepositoryObject
}

func (r *testRepository) Create(test *models.Test) error {
	return r.db.Create(test).Error
}

func (r *testRepository) Update(test *models.Test) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(test).Error
}

func (r *testRepository) Delete(test *models.Test) error {
	return r.db.Select(clause.Associations).Delete(test).Error
}

func (r *testRepository) Get(id uint) (*models.Test, error) {
	var testReq models.Test
	if err := r.db.
		Preload("Headers").
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
		r.db.Where(conditionStr, where...).Model(&testReq).Count(&total)
	} else {
		r.db.Model(&testReq).Count(&total)
	}

	if err := query.SetDB(r.db).
		SetPreloads("Headers", "RunTests", "TransportConfig", "TestGroup").
		SetModel(models.Test{}).
		SetWhere(conditionStr, where...).
		List(&testReq); err != nil {
		return nil, 0, err
	}

	return testReq, total, nil
}
