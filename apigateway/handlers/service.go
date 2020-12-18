package handlers

import (
	"context"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/s3f4/go-load/apigateway/library/log"
	res "github.com/s3f4/go-load/apigateway/library/response"
)

// ServiceHandler interface
type ServiceHandler interface {
	List(w http.ResponseWriter, r *http.Request)
}

type serviceHandler struct {
}

// NewServiceHandler returns new serviceHandler object
func NewServiceHandler() ServiceHandler {
	return &serviceHandler{}
}

func (h *serviceHandler) List(w http.ResponseWriter, r *http.Request) {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Debug(err)
		res.R500(w, r, err)
		return
	}

	//Retrieve a list of containers
	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		log.Debug(err)
		res.R500(w, r, err)
		return
	}

	res.R200(w, r, services)
}
