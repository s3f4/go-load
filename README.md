# Go Load

This project creates instances on [Digital Ocean](https://digitalocean.com) and install go load testing codes on them.

**Tech Stack**
 * [Terraform](https://terraform.io) for resource management,
 * [Ansible](https://docs.ansible.com/ansible/latest/index.html) for software provision,
 * [Go](https://golang.org) for backend,
 * [React.JS](https://reactjs.org) for frontend,
 * [Docker Swarm](https://docs.docker.com/engine/swarm/) for orchestration.

In order to create resources and install codes will be used the following codes.

```shell
cd infa/base
terraform init
terraform apply -var "do_token=abc.."
```



##### To delete all created resources

```shell
terraform destroy
```