// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// statsHandlersInterface is an autogenerated mock type for the statsHandlersInterface type
type statsHandlersInterface struct {
	mock.Mock
}

// List provides a mock function with given fields: w, r
func (_m *statsHandlersInterface) List(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}
