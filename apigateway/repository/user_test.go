package repository

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/stretchr/testify/assert"
)

func Test_User_Create(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewUserRepository(conn)
	user := &models.User{Email: "email", Password: "Password"}

	sqlMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`email`,`password`) VALUES (?,?)")).
		WithArgs(user.Email, user.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := r.Create(user)
	assert.Nil(t, err)
}

func Test_User_Get(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewUserRepository(conn)
	u := &models.User{ID: 1, Email: "email", Password: "Password"}

	rows := sqlmock.NewRows([]string{"id", "email", "password"}).
		AddRow(1, "email", "Password")

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`id` = ? ORDER BY `users`.`id` LIMIT 1")).
		WithArgs(u.ID).
		WillReturnRows(rows)

	user, err := r.Get(1)
	assert.Equal(t, user, u)
	assert.Nil(t, err)
}

func Test_User_Get_Error(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewUserRepository(conn)
	user := &models.User{ID: 1, Email: "email", Password: "Password"}

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`id` = ? ORDER BY `users`.`id` LIMIT 1")).
		WithArgs(user.ID).
		WillReturnError(errors.New(""))

	u, err := r.Get(1)
	assert.NotEqual(t, user, u)
	assert.NotNil(t, err)
}

func Test_User_GetByEmailAndPassword(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewUserRepository(conn)
	user := &models.User{ID: 1, Email: "email", Password: "Password"}

	rows := sqlmock.NewRows([]string{"id", "email", "password"}).
		AddRow(1, "email", "Password")

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE email=? AND password=? LIMIT 1")).
		WithArgs(user.Email, user.Password).
		WillReturnRows(rows)

	u, err := r.GetByEmailAndPassword(user)
	assert.Equal(t, user, u)
	assert.Nil(t, err)
}

func Test_User_GetByEmailAndPassword_Error(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewUserRepository(conn)
	user := &models.User{ID: 1, Email: "email", Password: "Password"}

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE email=? AND password=? LIMIT 1")).
		WithArgs(user.Email, user.Password).
		WillReturnError(errors.New(""))

	u, err := r.GetByEmailAndPassword(user)
	assert.NotEqual(t, user, u)
	assert.NotNil(t, err)
}

func Test_User_List(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewUserRepository(conn)
	allUsers := []*models.User{{ID: 1, Email: "email", Password: "Password"}}
	rows := sqlmock.NewRows([]string{"id", "email", "password"}).
		AddRow(1, "email", "Password")

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnRows(rows)

	users, err := r.List()
	assert.Equal(t, allUsers, users)
	assert.Nil(t, err)
}

func Test_User_List_Error(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewUserRepository(conn)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnError(errors.New(""))

	_, err := r.List()
	assert.NotNil(t, err)
}
