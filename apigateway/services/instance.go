package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/docker/docker/api/types/filters"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/library/log"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
	"github.com/s3f4/go-load/apigateway/template"
)

// InstanceService ...
type InstanceService interface {
	BuildTemplate(iReq models.InstanceConfig) (int, error)
	SpinUp() error
	Destroy() error
	ScaleWorkers(workerCount int) error
	Show(string) (string, error)
	ShowSwarmNodes() ([]swarm.Node, error)
	GetInstanceInfo() (*models.InstanceConfig, error)
	GetInstanceInfoFromTerraform() (string, error)
}

type instanceService struct {
	command    library.Command
	repository repository.InstanceRepository
}

// NewInstanceService returns an InstanceService object
func NewInstanceService(
	r repository.InstanceRepository,
	command library.Command,
) InstanceService {
	return &instanceService{
		repository: r,
		command:    command,
	}
}

func (s *instanceService) BuildTemplate(iReq models.InstanceConfig) (int, error) {
	f, err := os.OpenFile("./infra/workers.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
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
		"template/workers.tpl",
		"infra/workers.tf",
		map[string]interface{}{
			"Instances": instances,
			"env":       os.Getenv("APP_ENV"),
		},
	)

	if err := t.Write(); err != nil {
		return 0, err
	}

	if err := s.repository.Create(&iReq); err != nil {
		return 0, err
	}

	// returns how many worker will be created.
	return index, nil
}

// Spin Up instances
func (s *instanceService) SpinUp() error {
	if _, err := s.command.Run("cd infra;terraform apply -auto-approve;"); err != nil {
		log.Info(err)
		return err
	}

	s.command.Run("echo 'sleeping 20 secs for initializing'; sleep 20;")

	// don't try to join worker nodes to swarm
	// if env is development
	if os.Getenv("APP_ENV") != "development" {
		// install docker to worker nodes
		if err := s.installDockerToWNodes(); err != nil {
			log.Info(err)
			return err
		}

		if err := s.runAnsibleCommands(); err != nil {
			log.Info(err)
			return err
		}

		if err := s.joinWNodesToSwarm(); err != nil {
			log.Info(err)
			return err
		}

		if err := s.addLabels(); err != nil {
			log.Info(err)
			return err
		}
	}

	return nil
}

// installDockerToWNodes installs docker to worker nodes to join swarm
func (s *instanceService) installDockerToWNodes() error {
	output, err := s.command.Run(buildAnsibleCommand("docker-playbook.yml", ""))
	log.Info(string(output))
	return err
}

// Swarm nodes
func (s *instanceService) swarmInspect() (swarm.Swarm, error) {
	context := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Info(err)
		return swarm.Swarm{}, err
	}

	return cli.SwarmInspect(context)
}

func (s *instanceService) ScaleWorkers(workerCount int) error {
	context := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Error(err)
		return err
	}

	service, err := s.getService("go-load_worker")
	if err != nil {
		log.Error(err)
		return err
	}

	options := types.ServiceUpdateOptions{}
	wc := uint64(workerCount)

	service.Spec.Mode.Replicated.Replicas = &wc
	if _, err := cli.ServiceUpdate(context, service.ID, service.Version, service.Spec, options); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (s *instanceService) getService(name string) (*swarm.Service, error) {
	context := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Info(err)
		return nil, err
	}

	args := filters.NewArgs()
	args.Add("name", name)

	services, err := cli.ServiceList(context, types.ServiceListOptions{Filters: args})

	if services == nil || len(services) == 0 {
		return nil, errors.New("No matching service found for name " + name)
	}

	return &services[0], err
}

// joinWNodesToSwarm command runs ansible to join all workers to swarm
func (s *instanceService) joinWNodesToSwarm() error {
	swarm, err := s.swarmInspect()
	if err != nil {
		log.Info(err)
		return err
	}

	token := swarm.JoinTokens.Worker
	var addr string

	if os.Getenv("APP_ENV") != "development" {
		// If environment is production, read master ip address from inventory file
		addr, err = s.parseInventoryFile()
	} else {
		os.Getenv("APP_ENV")
		// If environment is not production read master ip address from ntppool site.
		addr, err = getIP()
	}

	if err != nil {
		log.Info(err)
		return nil
	}

	output, err := s.command.Run(buildAnsibleCommand("swarm-join.yml", fmt.Sprintf("--extra-vars '{\"token\":\"%s\",\"addr\": \"%s\"}'", token, addr)))
	log.Info(err)
	log.Info(string(output))
	return err
}

// runAnsibleCommands cert copies cert file to worker nodes to registry service
// hosts adds registry domain to /etc/hosts file
func (s *instanceService) runAnsibleCommands() error {
	output, err := s.command.Run(buildAnsibleCommand("cert.yml", ""))
	log.Debug(string(output))
	if err != nil {
		log.Info(err)
		return err
	}
	var addr string
	if os.Getenv("APP_ENV") != "development" {
		addr, err = s.parseInventoryFile()
	} else {
		addr, err = getIP()
	}

	if err != nil {
		log.Info(err)
		return err
	}

	output, err = s.command.Run(buildAnsibleCommand("hosts.yml", fmt.Sprintf("--extra-vars '{\"addr\": \"%s\"}'", addr)))
	if err != nil {
		log.Info(err)
		return err
	}

	return nil
}

// Destroy destroys worker instances
func (s *instanceService) Destroy() error {
	s.command.Run("cd infra;terraform destroy -auto-approve")
	s.command.Run("cd infra;rm -rf .terraform")
	s.command.Run("cd infra;rm -f terraform.tfstate*")
	t := template.NewInfraBuilder(
		"template/workers.tpl",
		"infra/workers.tf",
		map[string]interface{}{
			"Instances": nil,
			"env":       os.Getenv("APP_ENV"),
		},
	)

	if err := t.Write(); err != nil {
		log.Info(err)
		return err
	}

	s.command.Run("cd infra;terraform init;terraform apply -auto-approve")

	if err := s.repository.Delete(&models.InstanceConfig{}); err != nil {
		log.Info(err)
		return err
	}

	return nil
}

// Returns master's ip address
func (s *instanceService) parseInventoryFile() (string, error) {
	data, err := ioutil.ReadFile("./infra/ansible/inventory.tmpl")
	if err != nil {
		log.Info(err)
		return "", err
	}

	datas := strings.Split(string(data), "\n")
	return datas[1], err
}

// Show shows terraform outputs
func (s *instanceService) Show(output string) (string, error) {
	str := fmt.Sprintf("cd infra;terraform output -json %s", output)
	data, err := s.command.Run(str)
	log.Info(string(output))
	return string(data), err
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
func (s *instanceService) addLabels() error {
	context := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	var options types.NodeListOptions
	nodes, err := cli.NodeList(context, options)
	if err != nil {
		log.Info(err)
		return err
	}

	swarm, err := cli.SwarmInspect(context)
	if err != nil {
		log.Info(err)
		return err
	}

	// Loop all nodes.
	for _, node := range nodes {
		if strings.HasPrefix(node.Description.Hostname, "worker") {
			node.Spec.Annotations.Labels["role"] = "worker"
			if err := cli.NodeUpdate(context, node.ID, swarm.Version, node.Spec); err != nil {
				log.Info(err)
				return err
			}
		}
	}
	return nil
}

func (s *instanceService) GetInstanceInfo() (*models.InstanceConfig, error) {
	return s.repository.Get()
}

func (s *instanceService) GetInstanceInfoFromTerraform() (string, error) {
	output, err := s.command.Run("cd infra;terraform output -json workers")
	log.Info(string(output))
	return string(output), err
}

func buildAnsibleCommand(file string, extraVars string) string {
	c := fmt.Sprintf("cd ./infra/ansible; ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -i inventory.txt %s %s", file, extraVars)
	log.Info(c)
	return c
}

func getIP() (string, error) {
	resp, err := http.Get("https://www.mapper.ntppool.org/json")
	if err != nil {
		log.Info(err)
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Info(err)
		return "", err
	}

	ip := struct {
		HTTP string
		DNS  string
		EDNS string
	}{}

	if err = json.Unmarshal(body, &ip); err != nil {
		log.Info(err)
		return "", err
	}

	return ip.HTTP, nil
}
