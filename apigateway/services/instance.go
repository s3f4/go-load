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

	err = s.swarmInit()

	if err != nil {
		return err
	}
	return nil
}

func (s *instanceService) installDockerToWNodes() error {
	return mu.RunCommands("cd /etc/ansible && ansible-playbook -i inventory.txt docker-playbook.yml")
}

func (s *instanceService) swarmInit() error {
	masterIP, err := s.parseInventoryFile()
	if err != nil {
		return err
	}

	context := context.Background()
	cli, err := client.NewEnvClient()

	if err != nil {
		return err
	}

	req := swarm.InitRequest{
		AdvertiseAddr: masterIP,
	}

	_, err = cli.SwarmInit(context, req)

	if err != nil {
		return err
	}

	swarmIns, err := s.swarmInspect()
	if err != nil {
		return err
	}

	token := swarmIns.JoinTokens.Worker
	return s.joinWNodesToSwarm(token, masterIP)
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

// joinWNodesToSwarm command runs ansible to join all workers to swarm
func (s *instanceService) joinWNodesToSwarm(token, addr string) error {

	joinCommand := fmt.Sprintf(
		"ansible-playbook -i ./infra/ansible/inventory.txt ./infra/ansible/docker-swarm-init.yml "+
			"--extra-vars \"{token:%s,addr: %s}\"",
		token,
		addr,
	)

	mu.RunCommands(joinCommand)

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
