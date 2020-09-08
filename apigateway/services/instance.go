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
		iReq.InstanceCount,
	)

	if err := t.Write(); err != nil {
		return err
	}
	return nil
}

func (is *instanceService) SpinUp() error {
	executable := exec.Command("/bin/sh", "-c", "cd infra;terraform init;terraform apply")
	executable.Stdout = os.Stdout
	// terraformCmd := &exec.Cmd{
	// 	Path:   executable,
	// 	Stdout: os.Stdout,
	// 	Stdin:  os.Stdin,
	// 	Stderr: os.Stderr,
	// }

	if err := executable.Run(); err != nil {
		return err
	}
	return nil
}
func (is *instanceService) Run() error {
	return nil
}
func (is *instanceService) Destroy() error { return nil }
