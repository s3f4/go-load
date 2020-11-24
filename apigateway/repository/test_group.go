package repository

import (
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/gorm"
)

// TestGroupRepository ..
type TestGroupRepository interface {
	DB() *gorm.DB
	Create(*models.TestGroup) error
	Update(*models.TestGroup) error
	Delete(*models.TestGroup) error
	Get(id uint) (*models.TestGroup, error)
	List(*library.QueryBuilder) ([]models.TestGroup, int64, error)
}

type testGroupRepository struct {
	base BaseRepository
}

var testGroupRepositoryObject TestGroupRepository

// NewTestGroupRepository returns an testGroupRepository object
func NewTestGroupRepository() TestGroupRepository {
	if testGroupRepositoryObject == nil {
		return &testGroupRepository{
			base: NewBaseRepository(MYSQL),
		}
	}
	return testGroupRepositoryObject
}

func (r *testGroupRepository) DB() *gorm.DB {
	return r.base.GetDB()
}

func (r *testGroupRepository) Create(testGroup *models.TestGroup) error {
	return r.DB().Create(testGroup).Error
}

func (r *testGroupRepository) Update(testGroup *models.TestGroup) error {
	return r.DB().Model(testGroup).Updates(testGroup).Error
}

func (r *testGroupRepository) Delete(testGroup *models.TestGroup) error {
	return r.DB().Model(testGroup).Delete(testGroup).Error
}

func (r *testGroupRepository) Get(id uint) (*models.TestGroup, error) {
	var testReq models.TestGroup
	if err := r.DB().First(&testReq, id).Error; err != nil {
		return nil, err
	}
	return &testReq, nil
}

func (r *testGroupRepository) List(query *library.QueryBuilder) ([]models.TestGroup, int64, error) {
	var testReq []models.TestGroup
	var total int64
	r.DB().Model(&testReq).Count(&total)
	if err := query.
		SetDB(r.DB()).
		SetModel(models.TestGroup{}).
		SetPreloads("Tests", "Tests.Headers", "Tests.TransportConfig", "Tests.RunTests").
		List(&testReq); err != nil {
		return nil, 0, err
	}
	return testReq, total, nil
}
