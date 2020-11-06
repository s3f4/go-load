package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func parse(r *http.Request, model interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(model); err != nil {
		fmt.Println(err)
		return err
	}

	if err := validate.Struct(model); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
