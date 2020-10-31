package app

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/s3f4/go-load/apigateway/handlers"
	"github.com/s3f4/go-load/apigateway/middlewares"
)

func applyMiddlewares() {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	//router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	router.Use(cors.Handler)
}

// routeMap initializes routes.
func routeMap(*chi.Mux) {
	applyMiddlewares()

	router.Route("/user", func(router chi.Router) {
		router.Post("/login", handlers.UserHandler.Login)
		router.Get("/logout", handlers.UserHandler.Logout)
	})

	router.Post("/instances", handlers.InstanceHandler.SpinUp)
	router.Get("/instances", handlers.InstanceHandler.GetInstanceInfo)
	router.Get("/instances/regions", handlers.InstanceHandler.ShowRegions)
	router.Get("/instances/account", handlers.InstanceHandler.ShowAccount)
	router.Get("/instances/swarm-nodes", handlers.InstanceHandler.ShowSwarmNodes)
	router.Delete("/instances", handlers.InstanceHandler.Destroy)
	router.Get("/workers", handlers.WorkerHandler.List)
	router.Post("/workers", handlers.WorkerHandler.Stop)

	router.Post("/test_group/{ID}/start", handlers.TestGroupHandler.Start)
	router.Get("/test_group/{ID}", handlers.TestGroupHandler.List)
	router.Post("/test_group", handlers.TestGroupHandler.Create)
	router.Get("/test_group", handlers.TestGroupHandler.List)
	router.Put("/test_group", handlers.TestGroupHandler.Update)
	router.Delete("/test_group", handlers.TestGroupHandler.Delete)

	router.Route("/test", func(router chi.Router) {
		router.Post("/", handlers.TestHandler.Create)
		router.Get("/", handlers.TestHandler.List)

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
			router.Get("/stats", handlers.StatsHandler.List)
		})
		router.Get("/run_test", handlers.RunTestHandler.List)
	})
}
