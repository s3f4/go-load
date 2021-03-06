package middlewares

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/s3f4/go-load/apigateway/library"
	res "github.com/s3f4/go-load/apigateway/library/response"
)

// TestGroupCtx gets test with given id
func (m *Middleware) TestGroupCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		testGroupID, err := strconv.Atoi(chi.URLParam(r, "ID"))
		if err != nil {
			res.R404(w, r, library.ErrNotFound)
			return

		}
		testGroup, err := m.testGroupRepository.Get(uint(testGroupID))
		if err != nil {
			res.R404(w, r, library.ErrNotFound)
			return
		}
		ctx := context.WithValue(r.Context(), TestGroupCtxKey, testGroup)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
