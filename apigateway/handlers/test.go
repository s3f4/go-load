package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/services"
	. "github.com/s3f4/mu"
)

type testHandlerInterface interface {
	Insert(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Start(w http.ResponseWriter, r *http.Request)
}

type testHandler struct {
	service services.TestService
}

var (
	//TestHandler is handler.
	TestHandler testHandlerInterface = &testHandler{
		service: services.NewTestService(),
	}
)

func (h *testHandler) Insert(w http.ResponseWriter, r *http.Request) {
	var test models.Test
	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		fmt.Println(err)
		R400(w, "Bad Request")
		return
	}

	err := h.service.Insert(&test)
	if err != nil {
		fmt.Println(err)
		R500(w, err)
		return
	}
	R200(w, test)
}

func (h *testHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var test models.Test
	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		R400(w, "Bad Request")
		return
	}

	err := h.service.Delete(&test)
	if err != nil {
		R500(w, err)
		return
	}
	R200(w, test)
}

func (h *testHandler) Update(w http.ResponseWriter, r *http.Request) {
	var test models.Test
	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		R400(w, "Bad Request")
		return
	}

	err := h.service.Update(&test)
	if err != nil {
		R500(w, err)
		return
	}
	R200(w, test)
}

func (h *testHandler) Get(w http.ResponseWriter, r *http.Request) {
	var test models.Test
	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		R400(w, "Bad Request")
		return
	}

	tc, err := h.service.Get(&test)
	if err != nil {
		R500(w, err)
		return
	}
	R200(w, tc)
}
func (h *testHandler) List(w http.ResponseWriter, r *http.Request) {
	tests, err := h.service.List()
	if err != nil {
		R500(w, err)
		return
	}
	R200(w, tests)
}

func (h *testHandler) Start(w http.ResponseWriter, r *http.Request) {
	tests, err := h.service.List()
	if err != nil {
		R500(w, err)
		return
	}
	R200(w, tests)
}
