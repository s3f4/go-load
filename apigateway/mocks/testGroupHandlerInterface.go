// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// testGroupHandlerInterface is an autogenerated mock type for the testGroupHandlerInterface type
type testGroupHandlerInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: w, r
func (_m *testGroupHandlerInterface) Create(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Delete provides a mock function with given fields: w, r
func (_m *testGroupHandlerInterface) Delete(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Get provides a mock function with given fields: w, r
func (_m *testGroupHandlerInterface) Get(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// List provides a mock function with given fields: w, r
func (_m *testGroupHandlerInterface) List(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Update provides a mock function with given fields: w, r
func (_m *testGroupHandlerInterface) Update(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}
