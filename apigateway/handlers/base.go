package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/go-playground/validator/v10"
	"github.com/s3f4/go-load/apigateway/library/log"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func parse(r *http.Request, model interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(model); err != nil {
		log.Debug(err)
		return err
	}

	if err := validate.Struct(model); err != nil {
		log.Debug(err)
		return err
	}

	return nil
}

type handlerFunc func(w http.ResponseWriter, r *http.Request)

// makeRequest is used for http handlers tests.
func makeRequest(url, method string, handler handlerFunc, reader io.Reader) (*http.Response, []byte) {
	req := httptest.NewRequest(method, url, reader)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	return resp, body
}
