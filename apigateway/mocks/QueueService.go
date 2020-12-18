// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	services "github.com/s3f4/go-load/apigateway/services"
	mock "github.com/stretchr/testify/mock"
)

// QueueService is an autogenerated mock type for the QueueService type
type QueueService struct {
	mock.Mock
}

// Declare provides a mock function with given fields: queue
func (_m *QueueService) Declare(queue string) error {
	ret := _m.Called(queue)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(queue)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: queue
func (_m *QueueService) Delete(queue string) error {
	ret := _m.Called(queue)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(queue)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Listen provides a mock function with given fields: listenSpec
func (_m *QueueService) Listen(listenSpec *services.ListenSpec) {
	_m.Called(listenSpec)
}

// Send provides a mock function with given fields: queue, message
func (_m *QueueService) Send(queue string, message interface{}) error {
	ret := _m.Called(queue, message)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(queue, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
