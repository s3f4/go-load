package repository

import (
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/gorm"
)

// TestGroupRepository ..
type TestGroupRepository interface {
	Create(*models.TestGroup) error
	Update(*models.TestGroup) error
	Delete(*models.TestGroup) error
	Get(id uint) (*models.TestGroup, error)
	List(*library.QueryBuilder, string, ...interface{}) ([]models.TestGroup, int64, error)
}

type testGroupRepository struct {
	db *gorm.DB
}

var testGroupRepositoryObject TestGroupRepository

// NewTestGroupRepository returns an testGroupRepository object
func NewTestGroupRepository(db *gorm.DB) TestGroupRepository {
	if testGroupRepositoryObject == nil {
		testGroupRepositoryObject = &testGroupRepository{
			db: db,
		}
	}
	return testGroupRepositoryObject
}

func (r *testGroupRepository) Create(testGroup *models.TestGroup) error {
	return r.db.Create(testGroup).Error
}

func (r *testGroupRepository) Update(testGroup *models.TestGroup) error {
	return r.db.Model(testGroup).Updates(testGroup).Error
}

func (r *testGroupRepository) Delete(testGroup *models.TestGroup) error {
	return r.db.Model(testGroup).Delete(testGroup).Error
}

func (r *testGroupRepository) Get(id uint) (*models.TestGroup, error) {
	var testReq models.TestGroup
	if err := r.db.First(&testReq, id).Error; err != nil {
		return nil, err
	}
	return &testReq, nil
}

func (r *testGroupRepository) List(query *library.QueryBuilder, conditionStr string, where ...interface{}) ([]models.TestGroup, int64, error) {
	var testGroupReq []models.TestGroup
	var total int64
	if len(conditionStr) > 0 && where != nil {
		r.db.Where(conditionStr, where...).Model(&testGroupReq).Count(&total)
	} else {
		r.db.Model(&testGroupReq).Count(&total)
	}

	if err := query.
		SetDB(r.db).
		SetModel(models.TestGroup{}).
		SetPreloads("Tests", "Tests.Headers", "Tests.TransportConfig", "Tests.RunTests").
		SetWhere(conditionStr, where...).
		List(&testGroupReq); err != nil {
		return nil, 0, err
	}

	return testGroupReq, total, nil
}
