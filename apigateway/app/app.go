package app

import (
	"flag"
	"net/http"

	"github.com/go-chi/chi"
)

var router *chi.Mux

// Run apigateway Service
func Run() {
	router = chi.NewRouter()

	if err := runTemplates(); err != nil {
		return
	}

	initConnections()
	initHandlers()
	routeMap(router)
	migrate()

	port := flag.String("port", "3001", " default port is 3001")
	flag.Parse()

	if err := http.ListenAndServe(":"+*port, router); err != nil {
		panic(err)
	}
}
