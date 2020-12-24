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

func Test_Test_Create(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewTestRepository(conn)
	test := &models.Test{Name: "test"}

	sqlMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `tests` (`name`,`test_group_id`,`url`,`method`,`payload`,`request_count`,`goroutine_count`,`expected_response_code`,`expected_response_body`,`expected_first_byte_time`,`expected_connection_time`,`expected_dns_time`,`expected_tls_time`,`expected_total_time`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)")).
		WithArgs(
			test.Name,
			test.TestGroupID,
			test.URL,
			test.Method,
			test.Payload,
			test.RequestCount,
			test.GoroutineCount,
			test.ExpectedResponseCode,
			test.ExpectedResponseBody,
			test.ExpectedFirstByteTime,
			test.ExpectedConnectionTime,
			test.ExpectedDNSTime,
			test.ExpectedTLSTime,
			test.ExpectedTotalTime,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := r.Create(test)
	assert.Nil(t, err)
}

func Test_Test_Get(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewTestRepository(conn)
	test := &models.Test{
		ID:   1,
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
	}

	testRows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "test")

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `tests` WHERE `tests`.`id` = ? ORDER BY `tests`.`id` LIMIT 1")).
		WithArgs(test.ID).
		WillReturnRows(testRows)

	headerRows := sqlmock.NewRows([]string{"id", "test_id", "key", "value"}).
		AddRow(1, 1, "key", "value")

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `headers` WHERE `headers`.`test_id` = ?")).
		WithArgs(test.ID).
		WillReturnRows(headerRows)

	runTestRows := sqlmock.NewRows([]string{"id", "test_id"}).
		AddRow(1, 1)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `run_tests` WHERE `run_tests`.`test_id` = ?")).
		WithArgs(test.ID).
		WillReturnRows(runTestRows)

	transportConfigRows := sqlmock.NewRows([]string{"id", "test_id", "disable_keep_alives"}).
		AddRow(1, 1, true)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `transport_configs` WHERE `transport_configs`.`test_id` = ?")).
		WithArgs(test.ID).
		WillReturnRows(transportConfigRows)

	testResult, err := r.Get(test.ID)
	assert.Equal(t, test, testResult)
	assert.Nil(t, err)
}

func Test_Test_Get_Error(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewTestRepository(conn)
	test := &models.Test{ID: 1, Name: "test"}

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `tests` WHERE `tests`.`id` = ? ORDER BY `tests`.`id` LIMIT 1")).
		WithArgs(test.ID).
		WillReturnError(errors.New(""))

	testResult, err := r.Get(test.ID)
	assert.NotEqual(t, test, testResult)
	assert.NotNil(t, err)
}

func Test_Test_Update(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewTestRepository(conn)
	test := &models.Test{ID: 1, Name: "test"}

	sqlMock.ExpectExec(regexp.QuoteMeta("UPDATE `tests` SET `name`=? WHERE `id` = ?")).
		WithArgs(test.Name, test.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := r.Update(test)
	assert.Nil(t, err)
}
func Test_Test_Delete(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewTestRepository(conn)
	test := &models.Test{
		ID:   1,
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
		}}

	sqlMock.ExpectExec(regexp.QuoteMeta("DELETE FROM `run_tests` WHERE `run_tests`.`test_id` = ?")).
		WithArgs(test.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectExec(regexp.QuoteMeta("DELETE FROM `tests` WHERE `tests`.`id` = ?")).
		WithArgs(test.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectExec(regexp.QuoteMeta("DELETE FROM `transport_configs` WHERE `transport_configs`.`test_id` = ?")).
		WithArgs(test.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectExec(regexp.QuoteMeta("DELETE FROM `headers` WHERE `headers`.`test_id` = ?")).
		WithArgs(test.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := r.Delete(test)
	assert.Nil(t, err)
}

func Test_Test_List(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewTestRepository(conn)
	test := &models.Test{
		ID:   1,
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
	}

	testRows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "test")

	countRow := sqlmock.NewRows([]string{"count(1)"}).
		AddRow(1)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT count(1) FROM `tests`")).
		WillReturnRows(countRow)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `tests`")).
		WillReturnRows(testRows)

	headerRows := sqlmock.NewRows([]string{"id", "test_id", "key", "value"}).
		AddRow(1, 1, "key", "value")

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `headers` WHERE `headers`.`test_id` = ?")).
		WithArgs(test.ID).
		WillReturnRows(headerRows)

	runTestRows := sqlmock.NewRows([]string{"id", "test_id"}).
		AddRow(1, 1)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `run_tests` WHERE `run_tests`.`test_id` = ?")).
		WithArgs(test.ID).
		WillReturnRows(runTestRows)

	transportConfigRows := sqlmock.NewRows([]string{"id", "test_id", "disable_keep_alives"}).
		AddRow(1, 1, true)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `transport_configs` WHERE `transport_configs`.`test_id` = ?")).
		WithArgs(test.ID).
		WillReturnRows(transportConfigRows)

	tests, _, err := r.List(&library.QueryBuilder{}, "")
	assert.Nil(t, err)
	assert.Equal(t, *test, tests[0])
}

func Test_Test_List_Where(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewTestRepository(conn)
	test := &models.Test{ID: 1,
		Name: "test",
	}

	testRows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "test")

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `tests`")).
		WillReturnRows(testRows)

	headerRows := sqlmock.NewRows([]string{"id", "test_id", "key", "value"}).
		AddRow(1, 1, "key", "value")

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `headers` WHERE `headers`.`test_id` = ?")).
		WithArgs(test.ID).
		WillReturnRows(headerRows)

	runTestRows := sqlmock.NewRows([]string{"id", "test_id"}).
		AddRow(1, 1)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `run_tests` WHERE `run_tests`.`test_id` = ?")).
		WithArgs(test.ID).
		WillReturnRows(runTestRows)

	transportConfigRows := sqlmock.NewRows([]string{"id", "test_id", "disable_keep_alives"}).
		AddRow(1, 1, true)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `transport_configs` WHERE `transport_configs`.`test_id` = ?")).
		WithArgs(test.ID).
		WillReturnRows(transportConfigRows)

	_, _, err := r.List(&library.QueryBuilder{}, "name=?", "test")
	assert.Nil(t, err)
}

func Test_Test_List_Error(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewTestRepository(conn)

	countRow := sqlmock.NewRows([]string{"count(1)"}).
		AddRow(1)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT count(1) FROM `tests`")).
		WillReturnRows(countRow)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `tests`")).
		WillReturnError(errors.New(""))

	_, _, err := r.List(&library.QueryBuilder{}, "")
	assert.NotNil(t, err)
}
