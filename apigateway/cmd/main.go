package main

import (
	"log"

	"github.com/s3f4/go-load/apigateway"
	"github.com/s3f4/mu"
)

func main() {
	_, err := mu.RunCommands("cd infra;terraform init;terraform apply -auto-approve;")
	if err != nil {
		log.Panicf("error: %v", err)
	}
	apigateway.Run()
}
