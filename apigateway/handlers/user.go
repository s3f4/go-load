package handlers

import "net/http"

type userHandlerInterface interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type userHandler struct{}

var (
	// UserHandler is used for authentication
	UserHandler userHandlerInterface = &userHandler{}
)

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request)  {
	
}
func (h *userHandler) Logout(w http.ResponseWriter, r *http.Request) {}
