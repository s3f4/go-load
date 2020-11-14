package response

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

type handlerFunc func(w http.ResponseWriter, r *http.Request)

func sendReq(handler handlerFunc) (*http.Response, []byte) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	return resp, body
}

func TestR200(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		R200(w, r, "hello")
	}

	res, body := sendReq(handler)
	assert.Equal(t, res.StatusCode, http.StatusOK, "%d status is not equal to %d", res.StatusCode, http.StatusOK)

	var respBody Response
	err := json.Unmarshal(body, &respBody)

	assert.Nil(t, err, "response json unmarshal error")
	assert.Equal(t, respBody.Status, true, "Statuses are not equal")
	assert.Equal(t, "hello", respBody.Data)
}

func TestR200WithData(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		R200(w, r, map[string]interface{}{
			"a": "a",
			"b": "b",
		})
	}

	res, body := sendReq(handler)
	assert.Equal(t, res.StatusCode, http.StatusOK, "%d status is not equal to %d", res.StatusCode, http.StatusOK)

	var respBody Response
	err := json.Unmarshal(body, &respBody)
	assert.Nil(t, err, "response json unmarshal error")
	assert.Equal(t, respBody.Status, true, "Statuses are not equal")
	assert.Equal(t, map[string]interface{}{
		"a": "a",
		"b": "b",
	}, respBody.Data)
}

func TestR200WithStruct(t *testing.T) {
	type request struct {
		A string
		B bool
		C int64
	}

	req := request{
		A: "a",
		B: false,
		C: 55,
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		R200(w, r, req)
	}

	res, body := sendReq(handler)
	assert.Equal(t, res.StatusCode, http.StatusOK, "%d status is not equal to %d", res.StatusCode, http.StatusOK)

	var respBody Response
	err := json.Unmarshal(body, &respBody)
	assert.Nil(t, err, "response json unmarshal error")
	assert.Equal(t, respBody.Status, true, "Statuses are not equal")
	var newReq request
	mapstructure.Decode(respBody.Data, &newReq)
	assert.Equal(t, req, newReq)
}

func TestR400(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		R400(w, r, errors.New("hello"))
	}

	res, body := sendReq(handler)
	assert.Equalf(t, res.StatusCode, http.StatusBadRequest, "StatusCode: %d is not equal to %d", res.StatusCode, http.StatusBadRequest)

	var respBody Response
	err := json.Unmarshal(body, &respBody)

	assert.Nil(t, err, "response json unmarshal error")
	assert.Equal(t, respBody.Status, false, "Statuses are not equal")
	assert.Equal(t, "hello", respBody.Message)
}

func TestR401(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		R401(w, r, errors.New("hello"))
	}

	res, body := sendReq(handler)
	assert.Equal(t, res.StatusCode, http.StatusUnauthorized, "StatusCode: %d status is not equal to %d", res.StatusCode, http.StatusUnauthorized)

	var respBody Response
	err := json.Unmarshal(body, &respBody)

	assert.Nil(t, err, "response json unmarshal error")
	assert.Equal(t, respBody.Status, false, "Statuses are not equal")
	assert.Equal(t, "hello", respBody.Message)
}

//R403 returns Status forbidden
func TestR403(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		R403(w, r, errors.New("hello"))
	}

	res, body := sendReq(handler)
	assert.Equal(t, res.StatusCode, http.StatusForbidden, "StatusCode: %d status is not equal to %d", res.StatusCode, http.StatusForbidden)

	var respBody Response
	err := json.Unmarshal(body, &respBody)

	assert.Nil(t, err, "response json unmarshal error")
	assert.Equal(t, respBody.Status, false, "Statuses are not equal")
	assert.Equal(t, "hello", respBody.Message)
}

//R404 returns Status not found
func TestR404(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		R404(w, r, errors.New("hello"))
	}

	res, body := sendReq(handler)
	assert.Equal(t, res.StatusCode, http.StatusNotFound, "StatusCode: %d status is not equal to %d", res.StatusCode, http.StatusNotFound)

	var respBody Response
	err := json.Unmarshal(body, &respBody)

	assert.Nil(t, err, "response json unmarshal error")
	assert.Equal(t, respBody.Status, false, "Statuses are not equal")
	assert.Equal(t, "hello", respBody.Message)
}

//R422 returns Status unprocessable entity
func TestR422(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		R422(w, r, errors.New("hello"))
	}

	res, body := sendReq(handler)
	assert.Equal(t, res.StatusCode, http.StatusUnprocessableEntity, "StatusCode: %d status is not equal to %d", res.StatusCode, http.StatusUnprocessableEntity)

	var respBody Response
	err := json.Unmarshal(body, &respBody)

	assert.Nil(t, err, "response json unmarshal error")
	assert.Equal(t, respBody.Status, false, "Statuses are not equal")
	assert.Equal(t, "hello", respBody.Message)
}

//R500 returns Status internal server error
func TestR500(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		R500(w, r, errors.New("hello"))
	}

	res, body := sendReq(handler)
	assert.Equal(t, res.StatusCode, http.StatusInternalServerError, "StatusCode: %d status is not equal to %d", res.StatusCode, http.StatusInternalServerError)

	var respBody Response
	err := json.Unmarshal(body, &respBody)

	assert.Nil(t, err, "response json unmarshal error")
	assert.Equal(t, respBody.Status, false, "Statuses are not equal")
	assert.Equal(t, "hello", respBody.Message)
}
