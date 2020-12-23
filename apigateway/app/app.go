package app

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi"
)

var router *chi.Mux

// Run apigateway Service
func Run() {
	go Down()
	router = chi.NewRouter()
	initHandlers()
	routeMap(router)

	port := flag.String("port", "3001", " default port is 3001")
	flag.Parse()

	if err := http.ListenAndServe(":"+*port, router); err != nil {
		panic(err)
	}
}

//Down downs service when kill SIGINT came.
func Down() {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
	fmt.Println("\ni am dead")
	os.Exit(0)
}
