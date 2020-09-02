package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/s3f4/go-load/apigateway/models"
	template "github.com/s3f4/go-load/apigateway/template"
	. "github.com/s3f4/mu"
)

type instanceHandlerInterface interface {
	Init(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type instanceHandler struct{}

var (
	// InstanceHandler to use methods of handler.
	InstanceHandler instanceHandlerInterface = &instanceHandler{}
)

func (ih *instanceHandler) Init(w http.ResponseWriter, r *http.Request) {
	var instanceRequest models.InstanceRequest
	if err := json.NewDecoder(r.Body).Decode(&instanceRequest); err != nil {
		R400(w, err.Error())
		return
	}

	t := template.NewInfraBuilder(
		instanceRequest.Region,
		instanceRequest.InstanceSize,
		instanceRequest.InstanceCount,
	)

	if err := t.Write(); err != nil {
		R500(w, err.Error())
	}

	R200(w, map[string]interface{}{
		"data": instanceRequest,
	})
}

func (ih *instanceHandler) Destroy(w http.ResponseWriter, r *http.Request) {

}

func (ih *instanceHandler) List(w http.ResponseWriter, r *http.Request) {
	R200(w, Response{Data: map[string]interface{}{
		"ok": "ok",
	}})
}
