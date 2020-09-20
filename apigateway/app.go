package apigateway

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/s3f4/go-load/apigateway/handlers"
	"github.com/s3f4/go-load/apigateway/repository"
)

var router *chi.Mux

// Run apigateway Service
func Run() {
	go Down()
	router = chi.NewRouter()

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
	initHandlers()

	port := flag.String("port", "3001", " default port is 3001")
	flag.Parse()

	baseRepo := repository.NewBaseRepository(repository.MYSQL)
	baseRepo.Migrate()

	if err := http.ListenAndServe(":"+*port, router); err != nil {
		panic(err)
	}
}

func initHandlers() {
	router.Post("/instances", handlers.InstanceHandler.SpinUp)
	router.Get("/instances/regions", handlers.InstanceHandler.ShowRegions)
	router.Get("/instances/swarm-nodes", handlers.InstanceHandler.ShowSwarmNodes)
	router.Delete("/instances", handlers.InstanceHandler.Destroy)
	router.Get("/workers", handlers.WorkerHandler.List)
	router.Post("/workers", handlers.WorkerHandler.Stop)
	router.Post("/workers/run", handlers.WorkerHandler.Run)
	router.Get("/stats", handlers.StatsHandler.List)
}

//Down downs service when kill SIGINT came.
func Down() {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
	fmt.Println("\ni am dead")
	os.Exit(0)
}
