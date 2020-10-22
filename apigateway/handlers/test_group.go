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

type testGroupHandlerInterface interface {
	Start(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type testGroupHandler struct {
	service services.TestGroupService
}

var (
	//TestGroupHandler .
	TestGroupHandler testGroupHandlerInterface = &testGroupHandler{
		service: services.NewTestGroupService(),
	}
)

func (h *testGroupHandler) Start(w http.ResponseWriter, r *http.Request) {
	var run models.TestGroup
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

func (h *testGroupHandler) Create(w http.ResponseWriter, r *http.Request) {
	var testConfig models.TestGroup
	if err := json.NewDecoder(r.Body).Decode(&testConfig); err != nil {
		fmt.Println(err)
		R400(w, "Bad Request")
		return
	}

	fmt.Printf("%#v", testConfig)
	err := h.service.Create(&testConfig)
	if err != nil {
		fmt.Println(err)
		R500(w, err)
		return
	}
	R200(w, testConfig)
}

func (h *testGroupHandler) Update(w http.ResponseWriter, r *http.Request) {
	var testConfig models.TestGroup
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

func (h *testGroupHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var testConfig models.TestGroup
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

func (h *testGroupHandler) Get(w http.ResponseWriter, r *http.Request) {
	var testConfig models.TestGroup
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

func (h *testGroupHandler) List(w http.ResponseWriter, r *http.Request) {
	testConfig, err := h.service.List()
	if err != nil {
		R500(w, err)
		return
	}
	R200(w, testConfig)
}
