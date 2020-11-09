package middlewares

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/repository"
)

// TestCtx gets test with given id
func TestCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		testID, err := strconv.Atoi(chi.URLParam(r, "ID"))
		if err != nil {
			library.R400(w, r, err)
			return

		}
		tr := repository.NewTestRepository()
		test, err := tr.Get(uint(testID))
		if err != nil {
			library.R404(w, r, err)
			return
		}
		ctx := context.WithValue(r.Context(), TestCtxKey, test)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
