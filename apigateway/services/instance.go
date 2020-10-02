package services

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
	"github.com/s3f4/mu"
)

// InstanceService ...
type InstanceService interface {
	BuildTemplate(iReq models.InstanceConfig) error
	SpinUp() error
	Destroy() error
	ShowRegions() (string, error)
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
	_ = `
variable regions {
  default = [
    { index : 0, reg : "nyc1", instance_number : 1 },
    { index : 1, reg : "nyc1", instance_number : 2 },
    { index : 2, reg : "sgp1", instance_number : 1 },
    { index : 3, reg : "lon1", instance_number : 1 },
  ]
}

locals {
  regions = { for r in var.regions :
    r.index => r
  }
}

resource "digitalocean_droplet" "workers" {
  for_each = local.regions
  name     = "worker-${each.value.reg}-${each.value.instance_number}"
  region   = each.value.reg
  size     = "s-1vcpu-1gb"
  image    = "ubuntu-18-04-x64"
}
`

	for _, conf := range iReq.Configs {

		content := fmt.Sprintf(
			`resource "digitalocean_droplet" "worker_%s" {
			count  = %d
			image  = "%s"
			name   = "worker_%s_${count.index + 1}"

			region = "%s"
			size   = "%s"

			ssh_keys = [
					data.digitalocean_ssh_key.for_master.id
			]
		}`, conf.Region, conf.InstanceCount, conf.Image, conf.Region, conf.Region, conf.InstanceSize,
		)
		f.WriteString(content + "\n")
	}

	c, _ := ioutil.ReadAll(f)
	fmt.Println(string(c))
	fmt.Println("------")

	// iReq.Image = "ubuntu-18-04-x64"
	// iReq.InstanceSize = "s-1vcpu-1gb"

	// t := template.NewInfraBuilder(
	// 	iReq.Region,
	// 	iReq.InstanceSize,
	// 	iReq.Image,
	// 	iReq.InstanceCount,
	// )

	// if err := t.Write(); err != nil {
	// 	return err
	// }

	if err := s.repository.Insert(&iReq); err != nil {
		return err
	}

	return nil
}

// Spin Up instances
func (s *instanceService) SpinUp() error {
	_, err := mu.RunCommands("cd infra;terraform init;terraform apply -auto-approve")

	if err != nil {
		return err
	}
	return nil
}

// installDockerToWNodes installs docker to worker nodes to join swarm
func (s *instanceService) installDockerToWNodes() error {
	_, err := mu.RunCommands("cd ./infra/ansible; ansible-playbook -i inventory.txt docker-playbook.yml")
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
func (s *instanceService) joinWNodesToSwarm(token, addr string) error {

	joinCommand := fmt.Sprintf(
		"cd ./infra/ansible; ansible-playbook -i inventory.txt docker-swarm-init.yml "+
			"--extra-vars \"{token:%s,addr: %s}\"",
		token,
		addr,
	)

	mu.RunCommands(joinCommand)

	return nil
}

// Destroy destroys worker instances
func (s *instanceService) Destroy() error {
	mu.RunCommands("cd infra;terraform destroy -auto-approve")

	if err := os.Remove("./infra/terraform.tfstate"); err != nil && !os.IsNotExist(err) {
		return err
	}

	if err := os.Remove("./infra/terraform.tfstate.backup"); err != nil && !os.IsNotExist(err) {
		return err
	}

	// if err := os.Remove("./infra/workers.tf"); err != nil && !os.IsNotExist(err) {
	// 	return err
	// }

	if err := os.RemoveAll("./infra/.terraform"); err != nil && !os.IsNotExist(err) {
		return err
	}

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
	output1, err1 := mu.RunCommands("cd infra;export TF_LOG=true;terraform init;terraform apply -auto-approve;")
	fmt.Println(string(output1), err1)
	output, err := mu.RunCommands("cd infra;terraform output -json regions")
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

	swarm, err := cli.SwarmInspect(context)
	if err != nil {
		return err
	}

	for _, node := range nodes {
		if strings.HasPrefix(node.Description.Hostname, "worker") {
			node.Spec.Annotations.Labels["role"] = "worker"
			cli.NodeUpdate(context, nodes[0].ID, swarm.Version, nodes[0].Spec)
		}
	}
	return nil
}

func (s *instanceService) GetInstanceInfo() (*models.InstanceConfig, error) {
	return s.repository.Get()
}
