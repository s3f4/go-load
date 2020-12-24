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

func Test_Response_List(t *testing.T) {
	_, sqlMock, conn := ConnectMock(POSTGRES)
	r := NewResponseRepository(conn)

	responses := sqlmock.NewRows([]string{"id"}).
		AddRow(1)

	countRow := sqlmock.NewRows([]string{"count(1)"}).
		AddRow(1)

	sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT count(1) FROM "responses"`)).
		WillReturnRows(countRow)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"responses\"")).
		WillReturnRows(responses)

	_, _, err := r.List(&library.QueryBuilder{}, "")
	assert.Nil(t, err)
}

func Test_Response_List_Where(t *testing.T) {
	_, sqlMock, conn := ConnectMock(POSTGRES)
	r := NewResponseRepository(conn)
	runTest := &models.RunTest{
		ID: 1,
	}
	responses := sqlmock.NewRows([]string{"id"}).
		AddRow(1)
		
	countRow := sqlmock.NewRows([]string{"count(1)"}).
		AddRow(1)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"responses\"")).
		WillReturnRows(responses)

	sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "responses" WHERE run_test_id=?`)).
		WithArgs(runTest.ID).
		WillReturnRows(countRow)

	sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "responses" WHERE run_test_id=?`)).
		WithArgs(runTest.ID).
		WillReturnRows(responses)

	_, _, err := r.List(&library.QueryBuilder{}, "run_test_id=?", 1)
	assert.Nil(t, err)
}

func Test_Response_List_Error(t *testing.T) {
	_, sqlMock, conn := ConnectMock(POSTGRES)
	r := NewResponseRepository(conn)

	countRow := sqlmock.NewRows([]string{"count(1)"}).
		AddRow(1)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT count(1) FROM responses")).
		WillReturnRows(countRow)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM responses")).
		WillReturnError(errors.New(""))

	_, _, err := r.List(&library.QueryBuilder{}, "")
	assert.NotNil(t, err)
}
