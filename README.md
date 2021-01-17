# go-load

Go-load is a distributed load testing tool that can create droplets on [DigitalOcean](https://digitalocean.com) in different countries(regions) and create tests, make requests to servers/websites/apis, and measure request times such as first byte time, tls-dns connection time and it can test-compare response times, headers, body.


## Requirements

Go-load only works on macosx or linux computers. In order to use go-load you must [install terraform](https://www.terraform.io/downloads.html), and create an ssh key *(ssh key must be on ~/.ssh/id_rsa)* to be able to connect master instance and provide a digitalocean API key. 

```shell
ssh-keygen
 ```

Go to digitalocean's [API](https://cloud.digitalocean.com/account/api/tokens?i=8ebebf) section, take an API key

```shell
 echo 'do_token = "your api key"' > ./infra/base/terraform.tfvars 
 echo 'do_token = "your api key"' > ./apigateway/infra/terraform.tfvars 
 ```

 ## Install

Go to go-load folder on shell and run the following command.

```shell
make up 
```
This command spins up instances(droplets), uploads all codes to the master instance, installs docker, builds images and deploy images on docker swarm. This step will take a few minutes (~30 min), especially in the build images section. When the installing step is finished, you can get the master instance's IP address *(however, this is the address of the web interface)* with the following command.

```shell
make ip
```

Go to the master instance with the following command.
```shell
ssh root@{master instance IP}
```

Run the following command to see the services' status.

```shell
watch docker service ls
```

Wait until all services replicated except go-load_worker. If go-load_web and go-load_apigateway services are up and healty, you can go to the web interface with http://{master instance IP}, When you go, you will see the sign-in page, you can pass to the sign-up page and create a user with an email and password *(You can create only one user)*

## Usage

Create worker instances on the http://{master instance IP}/instances page. Then, go to /tests, create a test group, create a test and save the test group. Click the run button, this action makes requests and store the responses and measurement-comparing results. Go to the http://{master instance IP}/stats page, click the stats button of the test you created, and show the test results by clicking finished test button founds on the Finished Tests side.
## Remove 

Go to the http://{master instance IP}/instances page and click the destroy button to remove all worker instances, and remove the master and data instances with the following command: 

```shell
make destroy
```
This command will remove all resources created by go-load.

**Note**
*If you did not remove the worker instances on the instances page, you will have to remove all instances on digidalocean one by one.*

**Tech Stack**
 * [Terraform](https://terraform.io) for resource management,
 * [Ansible](https://docs.ansible.com/ansible/latest/index.html) for software provision,
 * [Go](https://golang.org) for backend,
 * [React](https://reactjs.org) for frontend,
 * [Docker Swarm](https://docs.docker.com/engine/swarm/) for orchestration.