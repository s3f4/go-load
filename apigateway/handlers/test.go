package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/services"
	. "github.com/s3f4/mu"
	"github.com/s3f4/mu/log"
)

type testHandlerInterface interface {
	Start(w http.ResponseWriter, r *http.Request)
	Insert(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	DeleteTest(w http.ResponseWriter, r *http.Request)
	UpdateTest(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
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

func (h *testHandler) Start(w http.ResponseWriter, r *http.Request) {
	var run models.TestConfig
	if err := json.NewDecoder(r.Body).Decode(&run); err != nil {
		R400(w, "Bad Request")
		return
	}

	if err := h.service.Start(&run); err != nil {
		log.Errorf("Worker Service Error: %s", err)
		R500(w, "worker service error")
		return
	}

	R200(w, "Test started.")
}

func (h *testHandler) Insert(w http.ResponseWriter, r *http.Request) {
	var testConfig models.TestConfig
	if err := json.NewDecoder(r.Body).Decode(&testConfig); err != nil {
		fmt.Println(err)
		R400(w, "Bad Request")
		return
	}

	err := h.service.Insert(&testConfig)
	if err != nil {
		fmt.Println(err)
		R500(w, err)
		return
	}
	R200(w, testConfig)
}

func (h *testHandler) Update(w http.ResponseWriter, r *http.Request) {
	var testConfig models.TestConfig
	if err := json.NewDecoder(r.Body).Decode(&testConfig); err != nil {
		R400(w, "Bad Request")
		return
	}
	err := h.service.Update(&testConfig)
	if err != nil {
		R500(w, err)
		return
	}
	R200(w, testConfig)
}

func (h *testHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var testConfig models.TestConfig
	if err := json.NewDecoder(r.Body).Decode(&testConfig); err != nil {
		R400(w, "Bad Request")
		return
	}

	err := h.service.Delete(&testConfig)
	if err != nil {
		R500(w, err)
		return
	}
	R200(w, testConfig)
}

func (h *testHandler) DeleteTest(w http.ResponseWriter, r *http.Request) {
	var test models.Test
	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		R400(w, "Bad Request")
		return
	}

	err := h.service.DeleteTest(&test)
	if err != nil {
		R500(w, err)
		return
	}
	R200(w, test)
}

func (h *testHandler) UpdateTest(w http.ResponseWriter, r *http.Request) {
	var test models.Test
	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		R400(w, "Bad Request")
		return
	}

	err := h.service.UpdateTest(&test)
	if err != nil {
		R500(w, err)
		return
	}
	R200(w, test)
}

func (h *testHandler) Get(w http.ResponseWriter, r *http.Request) {
	var testConfig models.TestConfig
	if err := json.NewDecoder(r.Body).Decode(&testConfig); err != nil {
		R400(w, "Bad Request")
		return
	}

	tc, err := h.service.Get(&testConfig)
	if err != nil {
		R500(w, err)
		return
	}
	R200(w, tc)
}
func (h *testHandler) List(w http.ResponseWriter, r *http.Request) {
	testConfig, err := h.service.List()
	if err != nil {
		R500(w, err)
		return
	}
	R200(w, testConfig)
}
