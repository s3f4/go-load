package services

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
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
	var err error
	if exists {
		err = mu.RunCommands("cd infra;terraform init;terraform apply -auto-approve")
	} else {
		err = mu.RunCommands("cd infra;terraform init;terraform apply -auto-approve")
	}

	if err != nil {
		return err
	}

	// err = s.swarmInit()
	s.ReadInventoryFile()
	s.swarmInspect()
	return err
}

func (s *instanceService) installDockerToWNodes() error {
	return mu.RunCommands("cd /etc/ansible && ansible-playbook -i inventory.txt docker-playbook.yml")
}

func (s *instanceService) swarmInit() error {
	context := context.Background()
	cli, err := client.NewEnvClient()

	if err != nil {
		panic(err)
	}

	req := swarm.InitRequest{
		ListenAddr: "eth0:2377",
	}

	_, err = cli.SwarmInit(context, req)

	if err != nil {
		return err
	}

	return nil
}

// Swarm nodes
func (s *instanceService) swarmInspect() (swarm.Swarm, error) {
	context := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	return cli.SwarmInspect(context)
}

func (s *instanceService) joinWNodesToSwarm() error {
	// todo ansible
	return nil
}

func (s *instanceService) Run() error {

	return nil
}

func (s *instanceService) Destroy() error {
	mu.RunCommands("cd infra;terraform destroy -auto-approve")
	return nil
}

// Returns master's ip address
func (s *instanceService) parseInventoryFile() (string, error) {
	data, err := ioutil.ReadFile("./infra/inventory.tmpl")
	if err != nil {
		return "", err
	}

	datas := strings.Split(string(data), "\n")
	return datas[1], err
}
