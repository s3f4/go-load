package middlewares

import (
	"context"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/services"
	"github.com/s3f4/mu/log"
)

// AuthCtx gets test with given id
func AuthCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ts := services.NewTokenService()
		as := services.NewAuthService()

		if r.Method == "GET" {
			w.Header().Set("X-CSRF-Token", csrf.Token(r))
		}

		if _, err := ts.VerifyToken(r, "at"); err != nil {
			log.Debug(err)
			library.R401(w, r, err)
			return
		}

		access, err := ts.GetDetailsFromToken(r, "at")
		if err != nil {
			log.Debug(err)
			library.R401(w, r, err)
			return
		}

		userID, err := as.GetAuthCache(access.UUID)
		if err != nil {
			log.Debug(err)
			library.R401(w, r, err)
			return
		}

		if r.RemoteAddr != access.RemoteAddr || r.UserAgent() != access.UserAgent {
			log.Debugf("r.RemoteAddr: %s", r.RemoteAddr)
			log.Debugf("access.RemoteAddr: %s", access.RemoteAddr)
			log.Debugf("r.UserAgent: %s", r.UserAgent())
			log.Debugf("access.UserAgent: %s", access.UserAgent)
			library.R401(w, r, err)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
