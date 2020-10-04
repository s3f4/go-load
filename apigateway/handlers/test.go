package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/services"
	. "github.com/s3f4/mu"
	"github.com/s3f4/mu/log"
)

type testHandlerInterface interface {
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

func (h *testHandler) Start(w http.ResponseWriter, r *http.Request) {
	var run models.TestConfig
	if err := json.NewDecoder(r.Body).Decode(&run); err != nil {
		R400(w, "Bad Request")
		return
	}

	if err := h.service.Start(run); err != nil {
		log.Errorf("Worker Service Error: %s", err)
		R500(w, "worker service error")
		return
	}

	R200(w, "Test started.")
}
