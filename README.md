# go-load

Go-load is a load testing library that can create droplets on [Digital Ocean](https://digitalocean.com) in different countries(regions) and create tests, make requests to servers/websites/apis and measure request times such as first byte time, tls-dns connection time.

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
This will take a few minutes (~30 min), especially in the build images section (currently go-load is building images and pushing to local private registry). When installing steps finished, you will see master ip address with

```shell
make ip
```
command. When you go http:// ip address, you will see signin page, you can pass to signup page and create a user with email and password (You can create only one user)
## Removing

First, you must remove your worker instances on the web interface and then, to remove master and data instances use the following command: 

```shell
make destroy
```



**Tech Stack**
 * [Terraform](https://terraform.io) for resource management,
 * [Ansible](https://docs.ansible.com/ansible/latest/index.html) for software provision,
 * [Go](https://golang.org) for backend,
 * [React](https://reactjs.org) for frontend,
 * [Docker Swarm](https://docs.docker.com/engine/swarm/) for orchestration.