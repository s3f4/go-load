SSH_FINGERPRINT=$(shell fp=`ssh-keygen -E md5 -lf ~/.ssh/id_rsa.pub | awk '{print $$2}'` && echo "$${fp/MD5:/}")

default:
	@echo "=============Building============="
	#docker build -f worker/Dockerfile.dev -t worker .

up: default
	@echo "=============Compose Up============="
	docker-compose -f docker-compose.dev.yml up -d  --build --remove-orphans

logs:
	docker-compose logs -f

down:
	docker-compose -f docker-compose.dev.yml down

test:
	go test -v -cover ./...

clean: down
	@echo "=============cleaning up============="
	rm -rf web/build
	#rm -rf web/node_modules
	rm -rf worker/cmd/worker
	rm -rf apigateway/cmd/apigateway
	rm -rf eventhandler/cmd/eventhandler
	docker system prune -f
	docker volume prune -f


create_ssh_for_master:
	ssh-keygen -t rsa -b 4096 -N '' -C "sefa@dehaa.com" -f ~/.ssh/id_rsa_for_master 

rm-files:
	rm -f apigateway/cmd/apigateway && \
	rm -f worker/cmd/worker && \
	rm -f eventhandler/cmd/eventhandler && \
	rm -rf web/node_modules && \
	rm -f infra/base/inventory.txt && \
	rm -f apigateway/log && \
	rm -f worker/log && \
	rm -f eventhandler/log  \

init :create_ssh_for_master
	cd infra/base && terraform init

apply:
	cd infra/base && export TF_LOG=true && terraform apply -auto-approve -var "public_key=$$HOME/.ssh/id_rsa.pub" \
  														 -var "private_key=$$HOME/.ssh/id_rsa" \
  														 -var "ssh_fingerprint=$(SSH_FINGERPRINT)" 
cpInventory:
	cd infra/base && cp inventory.txt ../../apigateway/infra/ansible/inventory.tmpl \
	&& echo "\n[workers]\n\$${workers}" >> ../../apigateway/infra/ansible/inventory.tmpl

up-instances: rm-files init apply cpInventory
	@echo "=============instances spinning up============="
	 
upload-inventory:
	cd infra/base && master=$$(terraform output master_ipv4_address) && scp inventory.txt root@$$master:/etc/ansible/inventory.txt

ansible-exec: upload-inventory
	cd infra/base && master=$$(terraform output master_ipv4_address) && ssh -t root@$$master 'cd /etc/ansible && ansible-playbook -i inventory.txt docker-playbook.yml'

ansible-ping: upload-inventory
	cd infra/base && master=$$(terraform output master_ipv4_address) && ssh -t root@$$master 'cd /etc/ansible && ansible all -i inventory.txt -m ping'

ssh-copy:
	@echo this command creates ssh key and copy the key other instances
	cd infra/base && master=$$(terraform output master_ipv4_address) && ssh -t root@$$master 'ssh-keygen' 

destroy:
	cd infra/base && terraform destroy -auto-approve -var "public_key=$$HOME/.ssh/id_rsa.pub" \
  														   -var "private_key=$$HOME/.ssh/id_rsa" \
  														   -var "ssh_fingerprint=$(SSH_FINGERPRINT)"

plan:
	cd infra/base && terraform plan -var "public_key=$$HOME/.ssh/id_rsa.pub" \
  														   -var "private_key=$$HOME/.ssh/id_rsa" \
  														   -var "ssh_fingerprint=$(SSH_FINGERPRINT)"

finger:
	@echo this is my fingerprint $(SSH_FINGERPRINT)

output:
	cd infra/base && terraform output regions

	