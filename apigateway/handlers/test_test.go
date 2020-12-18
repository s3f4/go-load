package handlers

import (
	"net/http"
	"strings"
	"testing"

	"github.com/s3f4/go-load/apigateway/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_TestCreate(t *testing.T) {
	service := new(mocks.TestService)
	repository := new(mocks.TestRepository)

	testHandler := NewTestHandler(service, repository)

	res, body := makeRequest("/test", http.MethodGet, testHandler.Create, strings.NewReader(`{"name":"test", "url":"url"}`))

	assert.Equal(t, res.Body, body)
	assert.Equal(t, res.StatusCode, http.StatusOK, "%d status is not equal to %d", res.StatusCode, http.StatusOK)
}
