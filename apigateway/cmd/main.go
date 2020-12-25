package main

import (
	"github.com/s3f4/go-load/apigateway/app"
	"github.com/s3f4/go-load/apigateway/library/log"
)

func main() {
	log.Info("apigateway is running....")
	app.Run()
}
