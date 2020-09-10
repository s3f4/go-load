package services

import (
	"fmt"

	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
	"github.com/s3f4/go-load/apigateway/template"
	"github.com/s3f4/mu"
)

// InstanceService ...
type InstanceService interface {
	BuildTemplate(iReq models.Instance) error
	SpinUp() error
	Run() error
	Destroy() error
}

type instanceService struct {
	repository repository.InstanceRepository
}

// NewInstanceService returns an InstanceService object
func NewInstanceService() InstanceService {
	return &instanceService{
		repository: repository.NewInstanceRepository(),
	}
}

func (s *instanceService) BuildTemplate(iReq models.Instance) error {
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
		fmt.Println(err)
		return err
	}

	if err := s.repository.Insert(&iReq); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *instanceService) SpinUp() error {
	exists := mu.DirExists("./infra/.terraform")
	if exists {
		mu.RunCommands("cd infra;terraform apply")
	} else {
		mu.RunCommands("cd infra;terraform init;terraform apply")
	}

	return nil
}

func (s *instanceService) addIpsToInventory() {
	mu.RunCommands("cd infra;terraform output")
}

func (s *instanceService) installDockerToWNodes() error {
	mu.RunCommands("cd /etc/ansible && ansible-playbook -i inventory.txt docker-playbook.yml")
	return nil
}

func (s *instanceService) joinWNodesToSwarm() error {
	// todo join docker swarm here
	return nil
}

func (s *instanceService) Run() error {

	return nil
}

func (s *instanceService) Destroy() error { return nil }
