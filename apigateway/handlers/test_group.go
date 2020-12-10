package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/library/log"
	res "github.com/s3f4/go-load/apigateway/library/response"
	"github.com/s3f4/go-load/apigateway/middlewares"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
	"gorm.io/gorm"
)

type testGroupHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type testGroupHandler struct {
	repository repository.TestGroupRepository
}

var (
	//TestGroupHandler .
	TestGroupHandler testGroupHandlerInterface = &testGroupHandler{
		repository: repository.NewTestGroupRepository(),
	}
)

func (h *testGroupHandler) Create(w http.ResponseWriter, r *http.Request) {
	var testConfig models.TestGroup
	if err := parse(r, &testConfig); err != nil {
		log.Debug(err)
		res.R400(w, r, library.ErrBadRequest)
		return
	}

	err := h.repository.Create(&testConfig)
	if err != nil {
		log.Debug(err)
		res.R500(w, r, library.ErrBadRequest)
		return
	}
	res.R200(w, r, testConfig)
}

func (h *testGroupHandler) Update(w http.ResponseWriter, r *http.Request) {
	var testConfig models.TestGroup
	if err := parse(r, &testConfig); err != nil {
		log.Debug(err)
		res.R400(w, r, library.ErrBadRequest)
		return
	}
	err := h.repository.Update(&testConfig)
	if err != nil {
		log.Debug(err)
		res.R500(w, r, err)
		return
	}
	res.R200(w, r, testConfig)
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
		res.R500(w, r, err)
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

	tc, err := h.repository.Get(testGroup.ID)
	if err != nil {
		log.Debug(err)
		res.R500(w, r, library.ErrInternalServerError)
		return
	}
	res.R200(w, r, tc)
}

func (h *testGroupHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query, ok := ctx.Value(middlewares.QueryCtxKey).(*library.QueryBuilder)
	if !ok {
		res.R422(w, r, library.ErrUnprocessableEntity)
		return
	}
	fmt.Printf("%#v", query)
	testConfig, total, err := h.repository.List(query)

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
		"data":  testConfig,
	})
}
