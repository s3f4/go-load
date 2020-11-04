package middlewares

import (
	"context"
	"net/http"

	"github.com/s3f4/go-load/apigateway/services"
	. "github.com/s3f4/mu"
)

// AuthCtx gets test with given id
func AuthCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		as := services.NewAuthService()

		if err := as.IsTokenValid(r); err != nil {
			R401(w, err)
			return
		}

		user := as.ExtractTokenMetadata(r)

		ctx := context.WithValue(r.Context(), RunTestCtxKey, runTest)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
