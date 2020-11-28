package app

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/gorilla/csrf"
	"github.com/s3f4/go-load/apigateway/handlers"
	res "github.com/s3f4/go-load/apigateway/library/response"
	"github.com/s3f4/go-load/apigateway/middlewares"
)

func applyMiddlewares() {
	secure := true
	if os.Getenv("APP_ENV") == "development" {
		secure = false
	}

	csrfMiddleware := csrf.Protect(
		[]byte(os.Getenv("CSRF_KEY")),
		csrf.TrustedOrigins([]string{"localhost:3000", "localhost:3001"}),
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
		router.Post("/signin", handlers.AuthHandler.Signin)
		router.Post("/signup", handlers.AuthHandler.Signup)
		router.Get("/signout", handlers.AuthHandler.Signout)
		router.Post("/_rt", handlers.AuthHandler.RefreshToken)
	})

	router.Group(func(router chi.Router) {
		router.Use(middlewares.AuthCtx)
		router.Get("/user/current_user", handlers.AuthHandler.CurrentUser)

		router.Route("/instances", func(router chi.Router) {
			router.Post("/", handlers.InstanceHandler.SpinUp)
			router.Get("/", handlers.InstanceHandler.GetInstanceInfo)
			router.Get("/terraform", handlers.InstanceHandler.GetInstanceInfoFromTerraform)
			router.Get("/regions", handlers.InstanceHandler.ShowRegions)
			router.Get("/account", handlers.InstanceHandler.ShowAccount)
			router.Get("/swarm-nodes", handlers.InstanceHandler.ShowSwarmNodes)
			router.Delete("/", handlers.InstanceHandler.Destroy)
		})

		router.Get("/workers", handlers.WorkerHandler.List)
		router.Post("/workers", handlers.WorkerHandler.Stop)
		router.Get("/services", handlers.ServiceHandler.List)

		router.Route("/test_group", func(router chi.Router) {
			router.Post("/", handlers.TestGroupHandler.Create)

			router.Route("/", func(router chi.Router) {
				router.Use(middlewares.QueryCtx)
				router.Get("/", handlers.TestGroupHandler.List)
			})

			router.Route("/{ID}", func(router chi.Router) {
				router.Use(middlewares.TestGroupCtx)
				router.Put("/", handlers.TestGroupHandler.Update)
				router.Delete("/", handlers.TestGroupHandler.Delete)
				router.Post("/start", handlers.TestGroupHandler.Start)
				router.Route("/tests", func(router chi.Router) {
					router.Use(middlewares.QueryCtx)
					router.Get("/", handlers.TestHandler.ListByTestGroupID)
				})
			})
		})

		router.Route("/test", func(router chi.Router) {
			router.Post("/", handlers.TestHandler.Create)
			router.Route("/", func(router chi.Router) {
				router.Use(middlewares.QueryCtx)
				router.Get("/", handlers.TestHandler.List)
			})

			router.Route("/{ID}", func(router chi.Router) {
				router.Use(middlewares.TestCtx)
				router.Post("/start", handlers.TestHandler.Start)
				router.Get("/", handlers.TestHandler.Get)
				router.Put("/", handlers.TestHandler.Update)
				router.Delete("/", handlers.TestHandler.Delete)
			})
		})

		router.Route("/run_test", func(router chi.Router) {
			router.Route("/{ID}", func(router chi.Router) {
				router.Use(middlewares.RunTestCtx)
				router.Get("/", handlers.RunTestHandler.Get)
				router.Delete("/", handlers.RunTestHandler.Delete)

				router.Route("/stats", func(router chi.Router) {
					router.Use(middlewares.QueryCtx)
					router.Get("/", handlers.StatsHandler.List)
				})
			})
			router.Get("/run_test", handlers.RunTestHandler.List)
		})
	})

}
