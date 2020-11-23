package middlewares

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	res "github.com/s3f4/go-load/apigateway/library/response"
	"github.com/s3f4/go-load/apigateway/repository"
)

// TestGroupCtx gets test with given id
func TestGroupCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		testGroupID, err := strconv.Atoi(chi.URLParam(r, "ID"))
		if err != nil {
			res.R400(w, r, err)
			return

		}
		tgr := repository.NewTestGroupRepository()
		testGroup, err := tgr.Get(uint(testGroupID))
		if err != nil {
			res.R404(w, r, err)
			return
		}
		ctx := context.WithValue(r.Context(), TestGroupCtxKey, testGroup)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
