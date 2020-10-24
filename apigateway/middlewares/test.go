package middlewares

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/s3f4/go-load/apigateway/repository"
	. "github.com/s3f4/mu"
)

// TestCtx gets test with given id
func TestCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		testID, err := strconv.Atoi(chi.URLParam(r, "ID"))
		if err != nil {
			R400(w, err)
			return

		}
		tr := repository.NewTestRepository()
		test, err := tr.Get(uint(testID))
		if err != nil {
			R404(w, err)
			return
		}
		ctx := context.WithValue(r.Context(), TestCtx, test)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
