// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	library "github.com/s3f4/go-load/apigateway/library"
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"

	models "github.com/s3f4/go-load/apigateway/models"
)

// TestRepository is an autogenerated mock type for the TestRepository type
type TestRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *TestRepository) Create(_a0 *models.Test) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Test) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB provides a mock function with given fields:
func (_m *TestRepository) DB() *gorm.DB {
	ret := _m.Called()

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: _a0
func (_m *TestRepository) Delete(_a0 *models.Test) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Test) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *TestRepository) Get(id uint) (*models.Test, error) {
	ret := _m.Called(id)

	var r0 *models.Test
	if rf, ok := ret.Get(0).(func(uint) *models.Test); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Test)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0, _a1, _a2
func (_m *TestRepository) List(_a0 *library.QueryBuilder, _a1 string, _a2 ...interface{}) ([]models.Test, int64, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _a2...)
	ret := _m.Called(_ca...)

	var r0 []models.Test
	if rf, ok := ret.Get(0).(func(*library.QueryBuilder, string, ...interface{}) []models.Test); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Test)
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

// Update provides a mock function with given fields: _a0
func (_m *TestRepository) Update(_a0 *models.Test) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Test) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
