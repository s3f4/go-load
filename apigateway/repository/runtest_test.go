package repository

import (
	"errors"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/stretchr/testify/assert"
)

func Test_RunTest_Create(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewRunTestRepository(conn)
	startTime := time.Now()
	endTime := time.Now()
	passed := true
	runTest := &models.RunTest{
		ID:        1,
		TestID:    1,
		StartTime: &startTime,
		EndTime:   &endTime,
		Passed:    &passed,
	}

	sqlMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `run_tests` (`test_id`,`start_time`,`end_time`,`passed`,`id`) VALUES (?,?,?,?,?)")).
		WithArgs(
			runTest.TestID,
			runTest.StartTime,
			runTest.EndTime,
			true,
			runTest.ID,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := r.Create(runTest)
	assert.Nil(t, err)
}

func Test_RunTest_Get(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewRunTestRepository(conn)
	runTest := &models.RunTest{
		ID: 1,
	}

	testRows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "test")

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `run_tests` WHERE `run_tests`.`id` = ? ORDER BY `run_tests`.`id` LIMIT 1")).
		WithArgs(runTest.ID).
		WillReturnRows(testRows)

	runTestResult, err := r.Get(runTest.ID)
	assert.Equal(t, runTest, runTestResult)
	assert.Nil(t, err)
}

func Test_RunTest_Get_Error(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewRunTestRepository(conn)
	runTest := &models.RunTest{
		ID: 1,
	}

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `run_tests` WHERE `run_tests`.`id` = ? ORDER BY `run_tests`.`id` LIMIT 1")).
		WithArgs(runTest.ID).
		WillReturnError(errors.New(""))

	_, err := r.Get(runTest.ID)
	assert.NotNil(t, err)
}

func Test_RunTest_Update(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewRunTestRepository(conn)
	startTime := time.Now()
	runTest := &models.RunTest{ID: 1, StartTime: &startTime}

	sqlMock.ExpectExec(regexp.QuoteMeta("UPDATE `run_tests` SET `start_time`=? WHERE `id` = ?")).
		WithArgs(runTest.StartTime, runTest.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := r.Update(runTest)
	assert.Nil(t, err)
}

func Test_RunTest_Delete(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewRunTestRepository(conn)
	runTest := &models.RunTest{ID: 1}

	sqlMock.ExpectExec(regexp.QuoteMeta("DELETE FROM `run_tests` WHERE `run_tests`.`id` = ?")).
		WithArgs(runTest.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := r.Delete(runTest)
	fmt.Println(err)
	assert.Nil(t, err)
}

func Test_RunTest_List(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewRunTestRepository(conn)
	runTest := &models.RunTest{
		ID: 1,
	}

	testRows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "test")

	countRow := sqlmock.NewRows([]string{"count(1)"}).
		AddRow(1)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT count(1) FROM `run_tests`")).
		WillReturnRows(countRow)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `run_tests`")).
		WillReturnRows(testRows)

	tests, _, err := r.List(&library.QueryBuilder{}, "")
	assert.Nil(t, err)
	assert.Equal(t, *runTest, tests[0])
}

func Test_RunTest_List_Where(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewRunTestRepository(conn)

	testRows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "test")

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `run_tests`")).
		WillReturnRows(testRows)

	_, _, err := r.List(&library.QueryBuilder{}, "name=?", "test")
	assert.Nil(t, err)
}

func Test_RunTest_List_Error(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewRunTestRepository(conn)

	countRow := sqlmock.NewRows([]string{"count(1)"}).
		AddRow(1)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT count(1) FROM `run_tests`")).
		WillReturnRows(countRow)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `run_tests`")).
		WillReturnError(errors.New(""))

	_, _, err := r.List(&library.QueryBuilder{}, "")
	assert.NotNil(t, err)
}
