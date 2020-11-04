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
		ts := services.NewTokenService()
		as := services.NewAuthService()

		if err := ts.IsTokenValid(r); err != nil {
			R401(w, err)
			return
		}

		access, err := ts.GetDetailsFromToken(r, "header")
		if err != nil {
			R401(w, err)
			return
		}

		userID, err := as.GetAuthCache(access)
		if err != nil {
			R401(w, err)
			return
		}

		ctx := context.WithValue(r.Context(), RunTestCtxKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
