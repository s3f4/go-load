package middlewares

import (
	"context"
	"net/http"
	"os"

	"github.com/s3f4/go-load/apigateway/library/log"
	res "github.com/s3f4/go-load/apigateway/library/response"
	"github.com/s3f4/go-load/apigateway/services"
)

// AuthCtx gets test with given id
func AuthCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ts := services.NewTokenService()
		as := services.NewAuthService()

		jwtToken, err := ts.VerifyToken(r, "at")

		if err != nil {
			log.Debug(err)
			res.R401(w, r, err)
			return
		}

		access, err := ts.GetDetailsFromToken(r, "at")
		if err != nil {
			log.Debug(err)
			res.R401(w, r, err)
			return
		}

		accessToken, err := as.GetAuthCache(access.UUID)
		if err != nil || accessToken != jwtToken.Raw {
			log.Debug(err)
			res.R401(w, r, err)
			return
		}

		if os.Getenv("APP_ENV") == "production" && len(os.Getenv("DOMAIN")) > 0 {
			if r.RemoteAddr != access.RemoteAddr || r.UserAgent() != access.UserAgent {
				log.Infof("r.RemoteAddr:%s\naccess.RemoteAddr:%s\n", r.RemoteAddr, access.RemoteAddr)
				log.Infof("r.UserAgent():%s\naccess.UserAgent:%s\n", r.UserAgent(), access.UserAgent)
				res.R401(w, r, err)
				return
			}
		}

		ctx := context.WithValue(r.Context(), UserIDKey, access.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
