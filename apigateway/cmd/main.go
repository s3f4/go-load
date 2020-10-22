package main

import (
	"os"

	"github.com/s3f4/go-load/apigateway/app"
	"github.com/s3f4/go-load/apigateway/template"
	"github.com/s3f4/mu"
	"github.com/s3f4/mu/log"
)

func sshTpl() {
	t := template.NewInfraBuilder(
		"template/data.tpl",
		"infra/data.tf",
		map[string]interface{}{"env": os.Getenv("APP_ENV")},
	)
	if err := t.Write(); err != nil {
		log.Errorf("error: ", err)
		return
	}
}

func main() {
	sshTpl()

	_, err := mu.RunCommands("cd infra;terraform init;terraform apply -auto-approve;")
	if err != nil {
		log.Errorf("error: %v", err)
		return
	}
	log.Info("apigateway is running....")
	app.Run()
}
