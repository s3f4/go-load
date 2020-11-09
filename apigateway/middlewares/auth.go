package middlewares

import (
	"context"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/services"
)

// AuthCtx gets test with given id
func AuthCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ts := services.NewTokenService()
		as := services.NewAuthService()

		if r.Method == "GET" {
			w.Header().Set("X-CSRF-Token", csrf.Token(r))
		}

		if err := ts.IsTokenValid(r); err != nil {
			library.R401(w, r, err)
			return
		}

		access, err := ts.GetDetailsFromToken(r, "header")
		if err != nil {
			library.R401(w, r, err)
			return
		}

		userID, err := as.GetAuthCache(access.UUID)
		if err != nil {
			library.R401(w, r, err)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
