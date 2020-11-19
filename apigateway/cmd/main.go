package main

import (
	"os"

	"github.com/s3f4/go-load/apigateway/app"
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/library/log"
	"github.com/s3f4/go-load/apigateway/template"
)

func main() {
	// Create data.tf file to create ssh key for workers.
	t := template.NewInfraBuilder(
		"template/data.tpl",
		"infra/data.tf",
		map[string]interface{}{"env": os.Getenv("APP_ENV")},
	)

	if err := t.Write(); err != nil {
		log.Errorf("error: ", err)
		return
	}

	if os.Getenv("APP_ENV") == "development" {
		_, _ = library.RunCommands("cd infra;terraform init;terraform destroy -auto-approve;")
	}

	_, err := library.RunCommands("cd infra;terraform init;terraform apply -auto-approve;")
	if err != nil {
		log.Errorf("error: %v", err)
		return
	}
	log.Info("apigateway is running....")
	app.Run()
}
