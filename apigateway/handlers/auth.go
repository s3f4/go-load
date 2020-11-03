package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
	"github.com/s3f4/go-load/apigateway/services"
	. "github.com/s3f4/mu"
)

type authHandlerInterface interface {
	Signin(w http.ResponseWriter, r *http.Request)
	Signout(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	ur repository.UserRepository
	as services.AuthService
}

var (
	// AuthHandler is used for authentication
	AuthHandler authHandlerInterface = &authHandler{
		ur: repository.NewUserRepository(),
		as: services.NewAuthService(),
	}
)

func (h *authHandler) Signin(w http.ResponseWriter, r *http.Request) {
	var userRequest models.User
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		R400(w, "Bad Request")
		return
	}

	user, err := h.ur.GetByEmailAndPassword(&userRequest)
	if err != nil {
		R404(w, "User Not Found")
		return
	}

}

func (h *authHandler) Signout(w http.ResponseWriter, r *http.Request) {

}
