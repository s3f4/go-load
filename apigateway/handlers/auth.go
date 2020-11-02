package handlers

import "net/http"

type authHandlerInterface interface {
	Signin(w http.ResponseWriter, r *http.Request)
	Signout(w http.ResponseWriter, r *http.Request)
}

type authHandler struct{}

var (
	// AuthHandler is used for authentication
	AuthHandler authHandlerInterface = &authHandler{}
)

func (h *authHandler) Signin(w http.ResponseWriter, r *http.Request) {

}

func (h *authHandler) Signout(w http.ResponseWriter, r *http.Request) {

}
