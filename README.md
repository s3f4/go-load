# go-load

This project creates instances on [Digital Ocean](https://digitalocean.com) and install go load testing codes on them.

**Tech Stack**
 * [Terraform](https://terraform.io) for resource management,
 * [Ansible](https://docs.ansible.com/ansible/latest/index.html) for software provision,
 * [Go](https://golang.org) for backend,
 * [React](https://reactjs.org) for frontend,
 * [Docker Swarm](https://docs.docker.com/engine/swarm/) for orchestration.

## Usage
If you want to use this tool you must provide a digitalocean apikey to create instances on digitalocean.

```shell
 echo 'do_token="your api key"' > ./infra/base/terraform.tfvars 
 echo 'do_token="your api key"' > ./apigateway/infra/terraform.tfvars 
 ```

In order to create resources and install codes, the following commands will be used.

```shell
make up 
```

this will take a few minutes

###### To install softwares on master and data instances
```shell
make swarm
```
###### To destroy all created resources

To be first, you must remove your worker instances on web interface and then to remove master and data instances use the following command: 

```shell
make destroy
```
