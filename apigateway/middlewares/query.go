package middlewares

import (
	"context"
	"net/http"

	"github.com/s3f4/go-load/apigateway/library"
)

// QueryCtx gets test with given id
func QueryCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := &library.QueryBuilder{}
		query.Build(r.URL.Query())
		ctx := context.WithValue(r.Context(), QueryCtxKey, query)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
