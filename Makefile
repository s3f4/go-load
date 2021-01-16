SSH_FINGERPRINT=$(shell fp=`ssh-keygen -E md5 -lf ~/.ssh/id_rsa.pub | awk '{print $$2}'` && echo "$${fp/MD5:/}")

default:
	@echo "=============Building============="
	#docker build -f worker/Dockerfile.dev -t worker .

create_ssh_for_master:
	ssh-keygen -t rsa -b 4096 -N '' -C "mail@example.com" -f ~/.ssh/id_rsa_for_master 

dev-logs:
	docker-compose logs -f

rm-files:
	rm -f apigateway/cmd/apigateway && \
	rm -rf apigateway/infra/.terraform && \
	rm -f apigateway/infra/ansible/inventory.tmpl && \
	rm -f apigateway/infra/terraform.tfstate* && \
	rm -f worker/cmd/worker && \
	rm -f eventhandler/cmd/eventhandler && \
	rm -rf web/node_modules && \
	rm -f infra/base/inventory.txt && \
	rm -rf apigateway/log && \
	rm -rf worker/log && \
	rm -rf eventhandler/log  && \
	rm -f ~/.ssh/id_rsa_for_master && \
	rm -f ~/.ssh/id_rsa_for_master.pub

gotest:
	cd apigateway && go test -v  -cover ./...

init :create_ssh_for_master
	cd infra/base && terraform init

apply:
	cd infra/base && export TF_LOG=true && terraform apply -auto-approve

up-dev: default
	@echo "=============Compose Up============="
	COMPOSE_HTTP_TIMEOUT=200 docker-compose -f docker-compose.yml up -d  --build --remove-orphans

down-dev:
	rm -rf apigateway/infra/.terraform && \
	rm -f apigateway/infra/.terraform* && \
	rm -f apigateway/infra/terraform.tfstate* && \
	docker-compose -f docker-compose.yml down

clean-dev: down-dev
	@echo "=============cleaning up============="
	rm -rf web/build
	#rm -rf web/node_modules
	rm -rf worker/cmd/worker
	rm -rf apigateway/cmd/apigateway
	rm -rf eventhandler/cmd/eventhandler
	docker system prune -f
	docker volume prune -f

up-debug: default
	@echo "=============Compose Up============="
	docker-compose -f docker-compose.debug.yml up -d  --build --remove-orphans

down-debug:
	rm -rf apigateway/infra/.terraform && \
	rm -f apigateway/infra/.terraform* && \
	rm -f apigateway/infra/terraform.tfstate* && \
	docker-compose -f docker-compose.debug.yml down

up-swarm-dev:
	@echo "=============Initializing local docker swarm============="
	docker swarm init 
	./build-dev.sh
	docker stack deploy -c swarm-dev.yml go-load

down-swarm-dev:
	docker swarm leave --force

cpInventory:
	cd infra/base && cp inventory.txt ../../apigateway/infra/ansible/inventory.tmpl \
	&& echo "\n[workers]\n\$${workers}" >> ../../apigateway/infra/ansible/inventory.tmpl

up-instances: rm-files init apply cpInventory
	@echo "=============instances spinning up============="
	cd infra/base && master=$$(terraform output -raw master_ipv4_address) && \
	ssh-keyscan -H $$master >> ~/.ssh/known_hosts 

upload-inventory:
	cd infra/base && master=$$(terraform output -raw master_ipv4_address) && scp inventory.txt root@$$master:/etc/ansible/inventory.txt && \
	scp ../../apigateway/infra/ansible/inventory.tmpl root@$$master:/root/app/apigateway/infra/ansible/inventory.tmpl

ansible-ping: 
	cd infra/base && master=$$(terraform output -raw master_ipv4_address) && ssh -t root@$$master 'cd /etc/ansible && ansible all -i inventory.txt -m ping'

swarm-prepare:
	cd infra/base && master=$$(terraform output -raw master_ipv4_address) && \
	ssh -t root@$$master "echo 'REACT_APP_API_BASE_URL=$$master' >> /root/app/web/.env && \
	cd /etc/ansible && \
	export ANSIBLE_HOST_KEY_CHECKING=False && \
	ansible-playbook -i inventory.txt known_hosts.yml && \
	ansible-playbook -i inventory.txt docker-playbook.yml && \
	ansible-playbook -i inventory.txt hosts.yml --extra-vars 'addr=$$master' && \
	ansible-playbook -i inventory.txt swarm-init-deploy.yml --extra-vars 'addr=$$master'" && \
	token=`ssh -t root@$$master -t docker swarm join-token worker -q` && \
	ssh -t root@$$master "cd /etc/ansible && \
	ansible-playbook -i inventory.txt swarm-join.yml --extra-vars 'token=$$token addr=$$master' && \
	ansible-playbook -i inventory.txt label.yml"

up: up-instances upload-inventory swarm-prepare
	
ssh-copy:
	@echo this command creates ssh key and copy the key other instances
	cd infra/base && master=$$(terraform output -raw master_ipv4_address) && ssh -t root@$$master 'ssh-keygen' 

destroy-terraform:
	cd infra/base && terraform destroy -auto-approve 

destroy: destroy-terraform rm-files
	
plan:
	cd infra/base && terraform plan 

finger:
	@echo this is my fingerprint $(SSH_FINGERPRINT)

output:
	cd infra/base && terraform output regions

ip:
	cd infra/base && terraform output -raw master_ipv4_address
	