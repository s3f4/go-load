package main

import (
	"os"
	"text/template"

	"github.com/s3f4/go-load/apigateway"
	"github.com/s3f4/mu"
	"github.com/s3f4/mu/log"
)

func sshTpl() {
	t, err := template.ParseFiles("infra/data.tpl")
	if err != nil {
		log.Panic(err)
		return
	}

	f, err := os.Create("infra/data.tf")
	if err != nil {
		log.Panicf("create file: ", err)
		return
	}
	defer f.Close()

	if err := t.Execute(f, map[string]string{"env": os.Getenv("APP_ENV")}); err != nil {
		log.Panicf("execute: ", err)
		return
	}
}

func main() {
	sshTpl()

	_, err := mu.RunCommands("cd infra;terraform init;terraform apply -auto-approve;")
	if err != nil {
		log.Panicf("error: %v", err)
	}
	apigateway.Run()
}
