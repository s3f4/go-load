package repository

import (
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/gorm"
)

// RunTestRepository ..
type RunTestRepository interface {
	Create(*models.RunTest) error
	Delete(*models.RunTest) error
	Update(*models.RunTest) error
	Get(id uint) (*models.RunTest, error)
	List(*library.QueryBuilder, string, ...interface{}) ([]models.RunTest, int64, error)
}

type runTestRepository struct {
	db *gorm.DB
}

// NewRunTestRepository returns an testRepository object
func NewRunTestRepository(db *gorm.DB) RunTestRepository {
	return &runTestRepository{
		db: db,
	}
}

func (r *runTestRepository) Create(runTest *models.RunTest) error {
	return r.db.Create(runTest).Error
}

func (r *runTestRepository) Delete(runTest *models.RunTest) error {
	return r.db.Model(runTest).Delete(runTest).Error
}

func (r *runTestRepository) Update(rt *models.RunTest) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(rt).Error
}

func (r *runTestRepository) Get(id uint) (*models.RunTest, error) {
	var testReq models.RunTest
	if err := r.db.First(&testReq, id).Error; err != nil {
		return nil, err
	}
	return &testReq, nil
}

func (r *runTestRepository) List(query *library.QueryBuilder, conditionStr string, where ...interface{}) ([]models.RunTest, int64, error) {
	var runTestReq []models.RunTest
	var total int64
	if len(conditionStr) > 0 && where != nil {
		r.db.Where(conditionStr, where...).Model(&runTestReq).Count(&total)
	} else {
		r.db.Model(&runTestReq).Count(&total)
	}

	if err := query.SetDB(r.db).
		SetModel(models.RunTest{}).
		SetWhere(conditionStr, where...).
		List(&runTestReq); err != nil {
		return nil, 0, err
	}

	return runTestReq, total, nil
}
