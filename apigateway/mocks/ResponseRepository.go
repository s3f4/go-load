// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	library "github.com/s3f4/go-load/apigateway/library"
	mock "github.com/stretchr/testify/mock"

	models "github.com/s3f4/go-load/apigateway/models"
)

// ResponseRepository is an autogenerated mock type for the ResponseRepository type
type ResponseRepository struct {
	mock.Mock
}

// List provides a mock function with given fields: _a0, _a1, _a2
func (_m *ResponseRepository) List(_a0 *library.QueryBuilder, _a1 string, _a2 ...interface{}) ([]models.Response, int64, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _a2...)
	ret := _m.Called(_ca...)

	var r0 []models.Response
	if rf, ok := ret.Get(0).(func(*library.QueryBuilder, string, ...interface{}) []models.Response); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Response)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(*library.QueryBuilder, string, ...interface{}) int64); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*library.QueryBuilder, string, ...interface{}) error); ok {
		r2 = rf(_a0, _a1, _a2...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
