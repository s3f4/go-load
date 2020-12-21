package middlewares

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/s3f4/go-load/apigateway/library"
	res "github.com/s3f4/go-load/apigateway/library/response"
)

// RunTestCtx gets test with given id
func (m *Middleware) RunTestCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		runTestID, err := strconv.Atoi(chi.URLParam(r, "ID"))
		if err != nil {
			res.R404(w, r, library.ErrNotFound)
			return
		}

		runTest, err := m.runTestRespository.Get(uint(runTestID))
		if err != nil {
			res.R404(w, r, library.ErrNotFound)
			return
		}
		ctx := context.WithValue(r.Context(), RunTestCtxKey, runTest)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
