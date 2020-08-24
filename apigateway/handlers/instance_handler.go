package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/s3f4/go-load/apigateway/models"
)

type instanceHandlerInterface interface {
	InitInstances(w http.ResponseWriter, r *http.Request)
	DestroyInstances(w http.ResponseWriter, r *http.Request)
	ListInstances(w http.ResponseWriter, r *http.Request)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

type instanceHandler struct{}

var (
	// InstanceHandler to use methods of handler.
	InstanceHandler instanceHandlerInterface = &instanceHandler{}
)

func (ih *instanceHandler) InitInstances(w http.ResponseWriter, r *http.Request) {
	var instanceRequest models.InstanceRequest
	if err := json.NewDecoder(r.Body).Decode(&instanceRequest); err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "JSON error")
	}
	respondWithJSON(w, http.StatusOK, instanceRequest)
}

func (ih *instanceHandler) DestroyInstances(w http.ResponseWriter, r *http.Request) {}

func (ih *instanceHandler) ListInstances(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, []byte("OK"))
}
