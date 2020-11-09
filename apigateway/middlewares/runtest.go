package middlewares

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/repository"
)

// RunTestCtx gets test with given id
func RunTestCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		runTestID, err := strconv.Atoi(chi.URLParam(r, "ID"))
		if err != nil {
			library.R400(w, r, err)
			return

		}
		rr := repository.NewRunTestRepository()
		runTest, err := rr.Get(uint(runTestID))
		if err != nil {
			library.R404(w, r, err)
			return
		}
		ctx := context.WithValue(r.Context(), RunTestCtxKey, runTest)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
