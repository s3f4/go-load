# go-load

Go-load is a distributed load testing tool that can create droplets on [DigitalOcean](https://digitalocean.com) in different countries(regions) and create tests, make requests to servers/websites/apis, and measure request times such as first byte time, tls-dns connection time and it can test-compare response times, headers, body.

## Usage
If you want to use this tool, you must provide a digitalocean apikey to create instances on digitalocean.

```shell
 echo 'do_token = "your api key"' > ./infra/base/terraform.tfvars 
 echo 'do_token = "your api key"' > ./apigateway/infra/terraform.tfvars 
 ```

In order to create resources and install codes, use the following command.

```shell
make up 
```
This step will take a few minutes (~30 min), especially in the build images section. When the installing step finished, you can get the master instance's IP address *(however, this is the address of the web interface)* with

```shell
make ip
```
command. Go to the master instance with the following command.
```shell
ssh root@your master instance IP
```

Write the following command to see the services' status.

```shell
watch docker service ls
```

Wait until all services replicated except go-load_worker. If go-load_web and go-load_apigateway services are up and healty, you can go to the web interface with http:// your master instance ip, When you go, you will see the signin page, you can pass to the signup page and create a user with an email and password *(You can create only one user)*
## Removing

Go to the */instances* page and destroy all worker instances, and remove the master and data instances with the following command: 

```shell
make destroy
```



**Tech Stack**
 * [Terraform](https://terraform.io) for resource management,
 * [Ansible](https://docs.ansible.com/ansible/latest/index.html) for software provision,
 * [Go](https://golang.org) for backend,
 * [React](https://reactjs.org) for frontend,
 * [Docker Swarm](https://docs.docker.com/engine/swarm/) for orchestration.