package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/apex/log"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/services"
	. "github.com/s3f4/mu"
	"go.uber.org/zap"
)

type instanceHandlerInterface interface {
	Init(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Run(w http.ResponseWriter, r *http.Request)
	ShowRegions(w http.ResponseWriter, r *http.Request)
	ShowSwarmNodes(w http.ResponseWriter, r *http.Request)
}

type instanceHandler struct {
	service services.InstanceService
}

var (
	// InstanceHandler to use methods of handler.
	InstanceHandler instanceHandlerInterface = &instanceHandler{
		services.NewInstanceService(),
	}
)

var logger *zap.Logger

func (h *instanceHandler) Init(w http.ResponseWriter, r *http.Request) {
	logger, _ = zap.NewProduction()
	log := logger.Sugar()
	var instanceRequest models.Instance
	if err := json.NewDecoder(r.Body).Decode(&instanceRequest); err != nil {
		log.Errorf("%s", err)
		R400(w, err.Error())
		return
	}

	if err := h.service.BuildTemplate(instanceRequest); err != nil {
		log.Errorf("%s", err)
		R500(w, err.Error())
		return
	}

	if err := h.service.SpinUp(); err != nil {
		log.Errorf("%s", err)
		R500(w, err.Error())
		return
	}

	R200(w, map[string]interface{}{
		"data": instanceRequest,
	})
}

func (h *instanceHandler) Destroy(w http.ResponseWriter, r *http.Request) {
	if err := h.service.Destroy(); err != nil {
		log.Errorf("%s", err)
		R500(w, err.Error())
		return
	}
}

func (h *instanceHandler) List(w http.ResponseWriter, r *http.Request) {
	R200(w, Response{Data: map[string]interface{}{
		"ok": "ok",
	}})
}

func (h *instanceHandler) Run(w http.ResponseWriter, r *http.Request) {

}

func (h *instanceHandler) ShowRegions(w http.ResponseWriter, r *http.Request) {
	output, err := h.service.ShowRegions()
	if err != nil {
		log.Errorf("%s", err)
		R500(w, err)
		return
	}
	R200(w, output)
}

func (h *instanceHandler) ShowSwarmNodes(w http.ResponseWriter, r *http.Request) {
	nodes, err := h.service.ShowSwarmNodes()
	if err != nil {
		log.Errorf("%s", err)
		R500(w, err)
		return
	}
	R200(w, nodes)
}
