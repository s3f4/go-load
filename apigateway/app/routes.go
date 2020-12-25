package app

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/gorilla/csrf"
	res "github.com/s3f4/go-load/apigateway/library/response"
)

func applyMiddlewares() {
	secure := true
	if os.Getenv("APP_ENV") == "development" {
		secure = false
	}

	csrfMiddleware := csrf.Protect(
		[]byte(os.Getenv("CSRF_KEY")),
		csrf.TrustedOrigins([]string{""localhost:3000", "localhost:3001""}),
		csrf.Secure(secure),
		csrf.Path("/"),
		csrf.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			resp, _ := json.Marshal(map[string]interface{}{
				"status":  false,
				"message": csrf.FailureReason(r).Error(),
				"headers": r.Header,
			})

			w.Write(resp)
		})),
	)

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	//router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:3001"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	router.Use(cors.Handler)
	router.Use(csrfMiddleware)
}

// routeMap initializes routes.
func routeMap(*chi.Mux) {
	applyMiddlewares()

	router.Get("/form", func(w http.ResponseWriter, r *http.Request) {
		res.R200(w, r, csrf.Token(r))
	})

	router.Route("/auth", func(router chi.Router) {
		router.Post("/signin", authHandler.Signin)
		router.Post("/signup", authHandler.Signup)
		router.Get("/signout", authHandler.Signout)
		router.Post("/_rt", authHandler.RefreshToken)
	})

	router.Group(func(router chi.Router) {
		router.Use(m.AuthCtx)
		router.Get("/user/current_user", authHandler.CurrentUser)

		router.Route("/instances", func(router chi.Router) {
			router.Post("/", instanceHandler.SpinUp)
			router.Get("/", instanceHandler.GetInstanceInfo)
			router.Get("/terraform", instanceHandler.GetInstanceInfoFromTerraform)
			router.Get("/regions", instanceHandler.ShowRegions)
			router.Get("/account", instanceHandler.ShowAccount)
			router.Get("/swarm-nodes", instanceHandler.ShowSwarmNodes)
			router.Delete("/", instanceHandler.Destroy)
		})

		router.Get("/workers", workerHandler.List)
		router.Post("/workers", workerHandler.Stop)
		router.Get("/services", serviceHandler.List)

		router.Route("/test_group", func(router chi.Router) {
			router.Post("/", testGroupHandler.Create)

			router.Route("/", func(router chi.Router) {
				router.Use(m.QueryCtx)
				router.Get("/", testGroupHandler.List)
			})

			router.Route("/{ID}", func(router chi.Router) {
				router.Use(m.TestGroupCtx)
				router.Put("/", testGroupHandler.Update)
				router.Delete("/", testGroupHandler.Delete)
				router.Route("/tests", func(router chi.Router) {
					router.Use(m.QueryCtx)
					router.Get("/", testHandler.ListByTestGroupID)
				})
			})
		})

		router.Route("/test", func(router chi.Router) {
			router.Post("/", testHandler.Create)
			router.Route("/", func(router chi.Router) {
				router.Use(m.QueryCtx)
				router.Get("/", testHandler.List)
			})

			router.Route("/{ID}", func(router chi.Router) {
				router.Use(m.TestCtx)
				router.Get("/", testHandler.Get)
				router.Put("/", testHandler.Update)
				router.Delete("/", testHandler.Delete)
				router.Route("/", func(router chi.Router) {
					router.Use(middleware.Timeout(20 * time.Second))
					router.Post("/start", testHandler.Start)
				})
				router.Route("/run_tests", func(router chi.Router) {
					router.Use(m.QueryCtx)
					router.Get("/", runTestHandler.ListByTestID)
				})
			})
		})

		router.Route("/run_test", func(router chi.Router) {
			router.Route("/{ID}", func(router chi.Router) {
				router.Use(m.RunTestCtx)
				router.Get("/", runTestHandler.Get)
				router.Delete("/", runTestHandler.Delete)

				router.Route("/stats", func(router chi.Router) {
					router.Use(m.QueryCtx)
					router.Get("/", statsHandler.List)
				})
			})
			router.Get("/", runTestHandler.List)
		})
	})

}
