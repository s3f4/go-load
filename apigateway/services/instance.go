package services

import (
	"os"
	"os/exec"

	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/template"
)

// InstanceServiceInterface ...
type InstanceServiceInterface interface {
	BuildTemplate(iReq models.Instance) error
	SpinUp() error
	Run() error
	Destroy() error
}

type instanceService struct {
}

var (
	// InstanceService handles all operations of instance handler
	InstanceService InstanceServiceInterface = &instanceService{}
)

func (is *instanceService) BuildTemplate(iReq models.Instance) error {
	// todo defaults...
	iReq.Region = "AMS3"
	iReq.InstanceOS = "ubuntu-18-04-x64"
	iReq.InstanceSize = "s-1vcpu-1gb"

	t := template.NewInfraBuilder(
		iReq.Region,
		iReq.InstanceSize,
		iReq.InstanceOS,
		iReq.InstanceCount,
	)

	if err := t.Write(); err != nil {
		return err
	}
	return nil
}

func (is *instanceService) SpinUp() error {
	exists := dirExists("./infra/.terraform")
	if exists {
		command("terraform apply")
	} else {
		command("terraform init;terraform apply")
	}

	return nil
}
func (is *instanceService) Run() error {
	return nil
}
func (is *instanceService) Destroy() error { return nil }

func command(command string) error {
	executable := exec.Command("/bin/sh", "-c", "cd infra;"+command)
	executable.Stdout = os.Stdout
	executable.Stderr = os.Stderr

	if err := executable.Start(); err != nil {
		return err
	}

	executable.Wait()
	return nil
}

func dirExists(dir string) bool {
	_, err := os.Stat(dir)
	return !os.IsNotExist(err)
}
