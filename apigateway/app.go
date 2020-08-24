package apigateway

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/s3f4/go-load/apigateway/handlers"
)

var router *mux.Router

// Run apigateway Service
func Run() {
	go Down()
	cors := gh.CORS(
		gh.AllowedHeaders([]string{"content-type"}),
		gh.AllowedOrigins([]string{"*"}),
		gh.AllowCredentials(),
	)
	router = mux.NewRouter().StrictSlash(true)
	initHandlers()

	router.Use(cors)

	port := flag.String("port", "3001", " default port is 3001")
	flag.Parse()

	if err := http.ListenAndServe(":"+*port, router); err != nil {
		panic(err)
	}
}

func initHandlers() {
	router.HandleFunc("/instances", handlers.InstanceHandler.Init).Methods("POST")
	router.HandleFunc("/instances", handlers.InstanceHandler.List).Methods("GET")
	router.HandleFunc("/instances", handlers.InstanceHandler.Destroy).Methods("DELETE")
	router.HandleFunc("/workers", handlers.WorkerHandler.List).Methods("GET")
	router.HandleFunc("/stats", handlers.StatsHandler.Get).Method("GET")
}

//Down downs service when kill SIGINT came.
func Down() {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
	fmt.Println("\ni am dead")
	os.Exit(0)
}
