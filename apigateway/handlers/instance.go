package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/s3f4/go-load/apigateway/library/log"
	res "github.com/s3f4/go-load/apigateway/library/response"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/services"
)

type instanceHandlerInterface interface {
	SpinUp(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
	ShowRegions(w http.ResponseWriter, r *http.Request)
	ShowAccount(w http.ResponseWriter, r *http.Request)
	ShowSwarmNodes(w http.ResponseWriter, r *http.Request)
	GetInstanceInfo(w http.ResponseWriter, r *http.Request)
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

func (h *instanceHandler) SpinUp(w http.ResponseWriter, r *http.Request) {
	var instanceConfig models.InstanceConfig
	if err := json.NewDecoder(r.Body).Decode(&instanceConfig); err != nil {
		log.Errorf(err.Error())
		res.R400(w, r, err.Error())
		return
	}

	if err := h.service.BuildTemplate(instanceConfig); err != nil {
		log.Errorf(err.Error())
		res.R500(w, r, err.Error())
		return
	}

	if err := h.service.SpinUp(); err != nil {
		log.Errorf(err.Error())
		res.R500(w, r, err.Error())
		return
	}

	res.R200(w, r, map[string]interface{}{
		"data": instanceConfig,
	})
}

func (h *instanceHandler) Destroy(w http.ResponseWriter, r *http.Request) {
	if err := h.service.Destroy(); err != nil {
		log.Errorf(err.Error())
		res.R500(w, r, err.Error())
		return
	}
}

func (h *instanceHandler) ShowRegions(w http.ResponseWriter, r *http.Request) {
	output, err := h.service.ShowRegions()
	if err != nil {
		log.Errorf(err.Error())
		res.R500(w, r, err)
		return
	}
	res.R200(w, r, output)
}

func (h *instanceHandler) ShowAccount(w http.ResponseWriter, r *http.Request) {
	output, err := h.service.ShowAccount()
	if err != nil {
		log.Errorf(err.Error())
		res.R500(w, r, err)
		return
	}
	res.R200(w, r, output)
}

func (h *instanceHandler) ShowSwarmNodes(w http.ResponseWriter, r *http.Request) {
	nodes, err := h.service.ShowSwarmNodes()
	if err != nil {
		log.Errorf(err.Error())
		res.R500(w, r, err)
		return
	}
	res.R200(w, r, nodes)
}

func (h *instanceHandler) GetInstanceInfo(w http.ResponseWriter, r *http.Request) {
	instanceConfig, err := h.service.GetInstanceInfo()
	if err != nil {
		res.R500(w, r, err)
		return
	}
	res.R200(w, r, instanceConfig)
}

func (h *instanceHandler) AddLabels(w http.ResponseWriter, r *http.Request) {
	err := h.service.AddLabels()
	if err != nil {
		log.Errorf(err.Error())
		res.R500(w, r, err)
		return
	}
	res.R200(w, r, "Labels has been added to worker nodes")
}
