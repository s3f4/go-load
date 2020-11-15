package handlers

import (
	"encoding/json"
	"net/http"

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
