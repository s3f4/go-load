package services

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
	"github.com/s3f4/go-load/apigateway/template"
	"github.com/s3f4/mu/log"
)

// InstanceService ...
type InstanceService interface {
	BuildTemplate(iReq models.InstanceConfig) error
	SpinUp() error
	Destroy() error
	ShowRegions() (string, error)
	ShowAccount() (string, error)
	ShowSwarmNodes() ([]swarm.Node, error)
	GetInstanceInfo() (*models.InstanceConfig, error)
	AddLabels() error
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

func (s *instanceService) BuildTemplate(iReq models.InstanceConfig) error {
	f, err := os.OpenFile("./infra/workers.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	index := 0
	var instances []string
	for _, conf := range iReq.Configs {
		for i := 1; i <= conf.Count; i++ {
			instances = append(instances, fmt.Sprintf("{ index : %d, reg : \"%s\", instance_number : %d },", index, conf.Region, i))
			index++
		}
	}

	t := template.NewInfraBuilder(
		instances,
	)

	if err := t.Write(); err != nil {
		return err
	}

	if err := s.repository.Insert(&iReq); err != nil {
		return err
	}

	return nil
}

// Spin Up instances
func (s *instanceService) SpinUp() error {
	if _, err := RunCommands("cd infra;terraform apply -auto-approve"); err != nil {

		return err
	}

	if err := s.runAnsibleCommands(); err != nil {
		return err
	}

	if err := s.installDockerToWNodes(); err != nil {
		return err
	}

	if err := s.joinWNodesToSwarm(); err != nil {
		return err
	}

	return nil
}

// installDockerToWNodes installs docker to worker nodes to join swarm
func (s *instanceService) installDockerToWNodes() error {
	output, err := RunCommands("cd ./infra/ansible; ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -i inventory.txt docker-playbook.yml")
	log.Debug(string(output))
	return err
}

// Swarm nodes
func (s *instanceService) swarmInspect() (swarm.Swarm, error) {
	context := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		return swarm.Swarm{}, err
	}

	return cli.SwarmInspect(context)
}

// joinWNodesToSwarm command runs ansible to join all workers to swarm
func (s *instanceService) joinWNodesToSwarm() error {

	swarm, err := s.swarmInspect()
	if err != nil {
		return err
	}

	token := swarm.JoinTokens.Worker
	addr, err := s.parseInventoryFile()

	if err != nil {
		return nil
	}

	joinCommand := fmt.Sprintf(
		"cd ./infra/ansible; ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -i inventory.txt swarm-join.yml "+
			"--extra-vars \"{token:%s,addr: %s}\"",
		token,
		addr,
	)

	output, err := RunCommands(joinCommand)
	log.Debug(string(output))
	return err
}

// runAnsibleCommands cert copies cert file to worker nodes to registry service
// hosts adds registry domain to /etc/hosts file
func (s *instanceService) runAnsibleCommands() error {
	output, err := RunCommands("cd ./infra/ansible; ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -i inventory.txt cert.yml")
	log.Debug(string(output))
	if err != nil {
		return err
	}

	addr, err := s.parseInventoryFile()
	if err != nil {
		return err
	}

	output, err = RunCommands("cd ./infra/ansible; ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -i inventory.txt hosts.yml" +
		fmt.Sprintf("--extra-vars \"{addr: %s}\"", addr))
	log.Debug(string(output))
	if err != nil {
		return err
	}

	output, err = RunCommands("cd ./infra/ansible; ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -i inventory.txt known_hosts.yml")
	log.Debug(string(output))
	if err != nil {
		return err
	}

	return nil
}

// Destroy destroys worker instances
func (s *instanceService) Destroy() error {
	RunCommands("cd infra;terraform destroy -auto-approve")
	RunCommands("cd infra;rm -rf .terraform")
	RunCommands("cd infra;rm -f terraform.tfstate*")
	t := template.NewInfraBuilder(
		nil,
	)

	if err := t.Write(); err != nil {
		return err
	}

	RunCommands("cd infra;terraform init;terraform apply -auto-approve")

	if err := s.repository.Delete(&models.InstanceConfig{}); err != nil {
		return err
	}

	return nil
}

// Returns master's ip address
func (s *instanceService) parseInventoryFile() (string, error) {
	data, err := ioutil.ReadFile("./infra/ansible/inventory.tmpl")
	if err != nil {
		return "", err
	}

	datas := strings.Split(string(data), "\n")
	return datas[1], err
}

// Terraform shows available regions
func (s *instanceService) ShowRegions() (string, error) {
	output, err := RunCommands("cd infra;terraform output -json regions")
	log.Debug(string(output))
	return string(output), err
}

// Terraform shows total droplet limit
func (s *instanceService) ShowAccount() (string, error) {
	output, err := RunCommands("cd infra;terraform output -json account")
	log.Debug(string(output))
	return string(output), err
}

// Shows swarm nodes
func (s *instanceService) ShowSwarmNodes() ([]swarm.Node, error) {
	context := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	var options types.NodeListOptions
	nodes, err := cli.NodeList(context, options)

	if err != nil {
		return nil, err
	}
	return nodes, nil
}

// Shows swarm nodes
func (s *instanceService) AddLabels() error {
	context := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	var options types.NodeListOptions
	nodes, err := cli.NodeList(context, options)
	if err != nil {
		return err
	}

	swarm, err := cli.SwarmInspect(context)
	if err != nil {
		return err
	}

	// Loop all nodes.
	for _, node := range nodes {
		if strings.HasPrefix(node.Description.Hostname, "worker") {
			node.Spec.Annotations.Labels["role"] = "worker"
			if err := cli.NodeUpdate(context, node.ID, swarm.Version, node.Spec); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *instanceService) GetInstanceInfo() (*models.InstanceConfig, error) {
	return s.repository.Get()
}

// RunCommands runs multiple commands
func RunCommands(command string) ([]byte, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(string(output))
		return output, err
	}
	fmt.Println(string(output))
	return output, nil
}
