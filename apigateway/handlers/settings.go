package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/s3f4/go-load/apigateway/library"
	res "github.com/s3f4/go-load/apigateway/library/response"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
)

// SettingsHandler interface
type SettingsHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type settingsHandler struct {
	repository repository.SettingsRepository
}

// NewSettingsHandler returns new settingsHandler object
func NewSettingsHandler(repository repository.SettingsRepository) SettingsHandler {
	return &settingsHandler{
		repository: repository,
	}
}

func (h *settingsHandler) Get(w http.ResponseWriter, r *http.Request) {
	settingParam := chi.URLParam(r, "setting")
	if len(settingParam) == 0 {
		res.R404(w, r, library.ErrNotFound)
		return
	}

	setting, err := h.repository.Get(models.SettingsKey(settingParam))
	if err != nil {
		res.R404(w, r, library.ErrNotFound)
		return
	}
	res.R200(w, r, setting)
}
