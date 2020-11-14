package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	res "github.com/s3f4/go-load/apigateway/library/response"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/services"
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
		res.R400(w, r, "Bad Request")
		return
	}

	if err := h.service.Start(&run); err != nil {
		log.Errorf("Worker Service Error: %s", err)
		res.R500(w, r, "worker service error")
		return
	}

	res.R200(w, r, "Test started.")
}

func (h *testGroupHandler) Create(w http.ResponseWriter, r *http.Request) {
	var testConfig models.TestGroup
	if err := json.NewDecoder(r.Body).Decode(&testConfig); err != nil {
		fmt.Println(err)
		res.R400(w, r, "Bad Request")
		return
	}

	fmt.Printf("%#v", testConfig)
	err := h.service.Create(&testConfig)
	if err != nil {
		fmt.Println(err)
		res.R500(w, r, err)
		return
	}
	res.R200(w, r, testConfig)
}

func (h *testGroupHandler) Update(w http.ResponseWriter, r *http.Request) {
	var testConfig models.TestGroup
	if err := json.NewDecoder(r.Body).Decode(&testConfig); err != nil {
		res.R400(w, r, "Bad Request")
		return
	}
	err := h.service.Update(&testConfig)
	if err != nil {
		res.R500(w, r, err)
		return
	}
	res.R200(w, r, testConfig)
}

func (h *testGroupHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var testConfig models.TestGroup
	if err := json.NewDecoder(r.Body).Decode(&testConfig); err != nil {
		res.R400(w, r, "Bad Request")
		return
	}

	err := h.service.Delete(&testConfig)
	if err != nil {
		res.R500(w, r, err)
		return
	}
	res.R200(w, r, testConfig)
}

func (h *testGroupHandler) Get(w http.ResponseWriter, r *http.Request) {
	var testConfig models.TestGroup
	if err := json.NewDecoder(r.Body).Decode(&testConfig); err != nil {
		res.R400(w, r, "Bad Request")
		return
	}

	tc, err := h.service.Get(&testConfig)
	if err != nil {
		res.R500(w, r, err)
		return
	}
	res.R200(w, r, tc)
}

func (h *testGroupHandler) List(w http.ResponseWriter, r *http.Request) {
	testConfig, err := h.service.List()
	if err != nil {
		res.R500(w, r, err)
		return
	}
	res.R200(w, r, testConfig)
}
