package app

import (
	"os"

	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/library/log"
	"github.com/s3f4/go-load/apigateway/template"
)

func runTemplates() error {
	// Create data.tf file to create ssh key for workers.
	t := template.NewInfraBuilder(
		"template/data.tpl",
		"infra/data.tf",
		map[string]interface{}{"env": os.Getenv("APP_ENV")},
	)

	if err := t.Write(); err != nil {
		log.Errorf("error: ", err)
		return err
	}

	if _, err := os.Stat("./infra/workers"); os.IsNotExist(err) {
		t := template.NewInfraBuilder(
			"template/workers.tpl",
			"infra/workers.tf",
			map[string]interface{}{"env": os.Getenv("APP_ENV")},
		)

		if err := t.Write(); err != nil {
			log.Errorf("error: ", err)
			return err
		}
	}

	command := library.NewCommand()
	if os.Getenv("APP_ENV") == "development" {
		command.Run("cd infra;terraform init;terraform destroy -auto-approve;")
	}

	if _, err := command.Run("cd infra;terraform init;terraform apply -auto-approve;"); err != nil {
		log.Errorf("error: %v", err)
		return err
	}
	return nil
}
