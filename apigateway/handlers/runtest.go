package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/s3f4/go-load/apigateway/middlewares"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/services"
	. "github.com/s3f4/mu"
)

type runTestHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type runTestHandler struct {
	service services.RunTestService
}

var (
	//RunTestHandler .
	RunTestHandler runTestHandlerInterface = &runTestHandler{
		service: services.NewRunTestService(),
	}
)

func (h *runTestHandler) Create(w http.ResponseWriter, r *http.Request) {
	var runTest models.RunTest
	if err := json.NewDecoder(r.Body).Decode(&runTest); err != nil {
		fmt.Println(err)
		R400(w, "Bad Request")
		return
	}

	err := h.service.Create(&runTest)
	if err != nil {
		fmt.Println(err)
		R500(w, err)
		return
	}
	R200(w, runTest)
}

func (h *runTestHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var runTest models.RunTest
	if err := json.NewDecoder(r.Body).Decode(&runTest); err != nil {
		R400(w, "Bad Request")
		return
	}

	err := h.service.Delete(&runTest)
	if err != nil {
		R500(w, err)
		return
	}
	R200(w, runTest)
}

func (h *runTestHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	runTest, ok := ctx.Value(middlewares.RunTestCtxKey).(*models.RunTest)
	if !ok {
		R422(w, "unprocessable entity")
		return
	}
	R200(w, runTest)
}
func (h *runTestHandler) List(w http.ResponseWriter, r *http.Request) {
	runTest, err := h.service.List()
	if err != nil {
		R500(w, err)
		return
	}
	R200(w, runTest)
}
