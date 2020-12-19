package handlers

import (
	"errors"
	"net/http"

	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/library/log"
	res "github.com/s3f4/go-load/apigateway/library/response"
	"github.com/s3f4/go-load/apigateway/middlewares"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
	"gorm.io/gorm"
)

// TestGroupHandler interface
type TestGroupHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type testGroupHandler struct {
	repository repository.TestGroupRepository
}

// NewTestGroupHandler returns new testGroupHandler object
func NewTestGroupHandler(repository repository.TestGroupRepository) TestGroupHandler {
	return &testGroupHandler{repository: repository}
}

func (h *testGroupHandler) Create(w http.ResponseWriter, r *http.Request) {
	var testGroup models.TestGroup
	if err := parse(r, &testGroup); err != nil {
		log.Debug(err)
		res.R400(w, r, library.ErrBadRequest)
		return
	}

	err := h.repository.Create(&testGroup)
	if err != nil {
		log.Debug(err)
		res.R500(w, r, library.ErrInternalServerError)
		return
	}
	res.R200(w, r, testGroup)
}

func (h *testGroupHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	testGroup, ok := ctx.Value(middlewares.TestGroupCtxKey).(*models.TestGroup)
	if !ok {
		res.R422(w, r, library.ErrUnprocessableEntity)
		return
	}

	var newTestGroup models.TestGroup
	newTestGroup.ID = testGroup.ID
	if err := parse(r, &newTestGroup); err != nil {
		log.Debug(err)
		res.R400(w, r, library.ErrBadRequest)
		return
	}
	err := h.repository.Update(&newTestGroup)
	if err != nil {
		log.Debug(err)
		res.R500(w, r, library.ErrInternalServerError)
		return
	}
	res.R200(w, r, newTestGroup)
}

func (h *testGroupHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	testGroup, ok := ctx.Value(middlewares.TestGroupCtxKey).(*models.TestGroup)
	if !ok {
		res.R422(w, r, library.ErrUnprocessableEntity)
		return
	}

	err := h.repository.Delete(testGroup)
	if err != nil {
		log.Debug(err)
		res.R500(w, r, library.ErrInternalServerError)
		return
	}
	res.R200(w, r, testGroup)
}

func (h *testGroupHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	testGroup, ok := ctx.Value(middlewares.TestGroupCtxKey).(*models.TestGroup)
	if !ok {
		res.R422(w, r, library.ErrUnprocessableEntity)
		return
	}

	res.R200(w, r, testGroup)
}

func (h *testGroupHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query, ok := ctx.Value(middlewares.QueryCtxKey).(*library.QueryBuilder)
	if !ok {
		res.R422(w, r, library.ErrUnprocessableEntity)
		return
	}
	testGroups, total, err := h.repository.List(query, "")

	if err != nil {
		log.Debug(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res.R404(w, r, library.ErrNotFound)
			return
		}
		res.R500(w, r, library.ErrInternalServerError)
		return
	}
	res.R200(w, r, map[string]interface{}{
		"total": total,
		"data":  testGroups,
	})
}
