package middlewares

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/s3f4/go-load/apigateway/repository"
	. "github.com/s3f4/mu"
)

// RunTestCtx gets test with given id
func RunTestCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		runTestID, err := strconv.Atoi(chi.URLParam(r, "ID"))
		if err != nil {
			R400(w, err)
			return

		}
		rr := repository.NewRunTestRepository()
		runTest, err := rr.Get(uint(runTestID))
		if err != nil {
			R404(w, err)
			return
		}
		ctx := context.WithValue(r.Context(), RunTestCtxKey, runTest)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
