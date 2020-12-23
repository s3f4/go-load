package repository

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/stretchr/testify/assert"
)

func Test_TestGroup_Create(t *testing.T) {
	_, sqlMock, conn := ConnectMock()
	r := NewTestGroupRepository(conn)
	testGroup := &models.TestGroup{Name: "test"}

	sqlMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `test_groups` (`name`) VALUES (?)")).
		WithArgs(
			testGroup.Name,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := r.Create(testGroup)
	assert.Nil(t, err)
}

func Test_TestGroup_Get(t *testing.T) {
	_, sqlMock, conn := ConnectMock()
	r := NewTestGroupRepository(conn)
	testGroup := &models.TestGroup{ID: 1, Name: "test"}

	testRows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "test")

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `test_groups` WHERE `test_groups`.`id` = ? ORDER BY `test_groups`.`id` LIMIT 1")).
		WithArgs(testGroup.ID).
		WillReturnRows(testRows)

	testResult, err := r.Get(testGroup.ID)
	assert.Equal(t, testGroup, testResult)
	assert.Nil(t, err)
}

func Test_TestGroup_Get_Error(t *testing.T) {
	_, sqlMock, conn := ConnectMock()
	r := NewTestGroupRepository(conn)
	testGroup := &models.TestGroup{ID: 1, Name: "test"}

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `test_groups` WHERE `test_groups`.`id` = ? ORDER BY `test_groups`.`id` LIMIT 1")).
		WithArgs(testGroup.ID).
		WillReturnError(errors.New(""))

	testResult, err := r.Get(testGroup.ID)
	assert.NotEqual(t, testGroup, testResult)
	assert.NotNil(t, err)
}

func Test_TestGroup_Update(t *testing.T) {
	_, sqlMock, conn := ConnectMock()
	r := NewTestGroupRepository(conn)
	testGroup := &models.TestGroup{ID: 1, Name: "test"}

	sqlMock.ExpectExec(regexp.QuoteMeta("UPDATE `test_groups` SET `name`=? WHERE `id` = ?")).
		WithArgs(testGroup.Name, testGroup.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := r.Update(testGroup)
	assert.Nil(t, err)
}
func Test_TestGroup_Delete(t *testing.T) {
	_, sqlMock, conn := ConnectMock()
	r := NewTestGroupRepository(conn)
	testGroup := &models.TestGroup{
		ID:   1,
		Name: "test",
	}

	sqlMock.ExpectExec(regexp.QuoteMeta("DELETE FROM `test_groups` WHERE `test_groups`.`id` = ?")).
		WithArgs(testGroup.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := r.Delete(testGroup)
	assert.Nil(t, err)
}

func Test_TestGroup_List(t *testing.T) {
	_, sqlMock, conn := ConnectMock()
	r := NewTestGroupRepository(conn)
	testGroup := &models.TestGroup{ID: 1,
		Name: "test",
		Tests: []*models.Test{
			{
				ID:          1,
				Name:        "test",
				TestGroupID: 1,
				TransportConfig: models.TransportConfig{
					TestID:            1,
					DisableKeepAlives: true,
				},
				Headers: []*models.Header{
					{ID: 1, TestID: 1, Key: "key", Value: "value"},
				},
				RunTests: []*models.RunTest{
					{ID: 1, TestID: 1},
				},
			},
		},
	}

	testGroupRows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "test")

	countRow := sqlmock.NewRows([]string{"count(1)"}).
		AddRow(1)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT count(1) FROM `test_groups`")).
		WillReturnRows(countRow)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `test_groups`")).
		WillReturnRows(testGroupRows)

	testRows := sqlmock.NewRows([]string{"id", "name", "test_group_id"}).
		AddRow(1, "test", 1)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `tests` WHERE `tests`.`test_group_id` = ?")).
		WithArgs(testGroup.ID).
		WillReturnRows(testRows)

	headerRows := sqlmock.NewRows([]string{"id", "test_id", "key", "value"}).
		AddRow(1, 1, "key", "value")

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `headers` WHERE `headers`.`test_id` = ?")).
		WithArgs(testGroup.Tests[0].ID).
		WillReturnRows(headerRows)

	runTestRows := sqlmock.NewRows([]string{"id", "test_id"}).
		AddRow(1, 1)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `run_tests` WHERE `run_tests`.`test_id` = ?")).
		WithArgs(testGroup.Tests[0].ID).
		WillReturnRows(runTestRows)

	transportConfigRows := sqlmock.NewRows([]string{"id", "test_id", "disable_keep_alives"}).
		AddRow(1, 1, true)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `transport_configs` WHERE `transport_configs`.`test_id` = ?")).
		WithArgs(testGroup.Tests[0].ID).
		WillReturnRows(transportConfigRows)

	testGroups, count, err := r.List(&library.QueryBuilder{}, "")
	assert.Equal(t, count, int64(1))
	assert.Nil(t, err)
	assert.Equal(t, *testGroup, testGroups[0])
}

func Test_TestGroup_List_Where(t *testing.T) {
	_, sqlMock, conn := ConnectMock()
	r := NewTestGroupRepository(conn)
	testGroup := &models.TestGroup{ID: 1,
		Name: "test",
		Tests: []*models.Test{
			{
				ID:   1,
				TestGroupID: 1,
				Name: "test",
				TransportConfig: models.TransportConfig{
					TestID:            1,
					DisableKeepAlives: true,
				},
				Headers: []*models.Header{
					{ID: 1, TestID: 1, Key: "key", Value: "value"},
				},
				RunTests: []*models.RunTest{
					{ID: 1, TestID: 1},
				},
			},
		},
	}

	testGroupRows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "test")

	countRow := sqlmock.NewRows([]string{"count(1)"}).
		AddRow(1)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT count(1) FROM `test_groups`")).
		WillReturnRows(countRow)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `test_groups`")).
		WillReturnRows(testGroupRows)

	testRows := sqlmock.NewRows([]string{"id", "name", "test_group_id"}).
		AddRow(1, "test", 1)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `tests` WHERE `tests`.`test_group_id` = ?")).
		WithArgs(testGroup.ID).
		WillReturnRows(testRows)

	headerRows := sqlmock.NewRows([]string{"id", "test_id", "key", "value"}).
		AddRow(1, 1, "key", "value")

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `headers` WHERE `headers`.`test_id` = ?")).
		WithArgs(testGroup.Tests[0].ID).
		WillReturnRows(headerRows)

	runTestRows := sqlmock.NewRows([]string{"id", "test_id"}).
		AddRow(1, 1)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `run_tests` WHERE `run_tests`.`test_id` = ?")).
		WithArgs(testGroup.Tests[0].ID).
		WillReturnRows(runTestRows)

	transportConfigRows := sqlmock.NewRows([]string{"id", "test_id", "disable_keep_alives"}).
		AddRow(1, 1, true)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `transport_configs` WHERE `transport_configs`.`test_id` = ?")).
		WithArgs(testGroup.Tests[0].ID).
		WillReturnRows(transportConfigRows)

	testGroups, count, err := r.List(&library.QueryBuilder{}, "name=?", "test")
	assert.NotNil(t, testGroups)
	assert.Equal(t, count, int64(1))
	assert.Nil(t, err)
	assert.Equal(t, *testGroup, testGroups[0])
}

func Test_TestGroup_List_Error(t *testing.T) {
	_, sqlMock, conn := ConnectMock()
	r := NewTestGroupRepository(conn)

	countRow := sqlmock.NewRows([]string{"count(1)"}).
		AddRow(1)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT count(1) FROM `test_groups`")).
		WillReturnRows(countRow)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `test_groups`")).
		WillReturnError(errors.New(""))

	testGroups, count, err := r.List(&library.QueryBuilder{}, "")
	assert.Nil(t, testGroups)
	assert.Equal(t, count, int64(0))
	assert.NotNil(t, err)
}
