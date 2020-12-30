// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	models "github.com/s3f4/go-load/apigateway/models"
	mock "github.com/stretchr/testify/mock"

	swarm "github.com/docker/docker/api/types/swarm"
)

// InstanceService is an autogenerated mock type for the InstanceService type
type InstanceService struct {
	mock.Mock
}

// BuildTemplate provides a mock function with given fields: iReq
func (_m *InstanceService) BuildTemplate(iReq models.InstanceConfig) (int, error) {
	ret := _m.Called(iReq)

	var r0 int
	if rf, ok := ret.Get(0).(func(models.InstanceConfig) int); ok {
		r0 = rf(iReq)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.InstanceConfig) error); ok {
		r1 = rf(iReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Destroy provides a mock function with given fields:
func (_m *InstanceService) Destroy() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetInstanceInfo provides a mock function with given fields:
func (_m *InstanceService) GetInstanceInfo() (*models.InstanceConfig, error) {
	ret := _m.Called()

	var r0 *models.InstanceConfig
	if rf, ok := ret.Get(0).(func() *models.InstanceConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.InstanceConfig)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInstanceInfoFromTerraform provides a mock function with given fields:
func (_m *InstanceService) GetInstanceInfoFromTerraform() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScaleWorkers provides a mock function with given fields: workerCount
func (_m *InstanceService) ScaleWorkers(workerCount int) error {
	ret := _m.Called(workerCount)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(workerCount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Show provides a mock function with given fields: _a0
func (_m *InstanceService) Show(_a0 string) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowSwarmNodes provides a mock function with given fields:
func (_m *InstanceService) ShowSwarmNodes() ([]swarm.Node, error) {
	ret := _m.Called()

	var r0 []swarm.Node
	if rf, ok := ret.Get(0).(func() []swarm.Node); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]swarm.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SpinUp provides a mock function with given fields:
func (_m *InstanceService) SpinUp() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
