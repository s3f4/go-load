SSH_FINGERPRINT=$(shell fp=`ssh-keygen -E md5 -lf ~/.ssh/id_rsa.pub | awk '{print $$2}'` && echo "$${fp/MD5:/}")

default:
	@echo "=============Building============="
	#docker build -f worker/Dockerfile.dev -t worker .

up: default
	@echo "=============Compose Up============="
	docker-compose -f docker-compose.dev.yml up -d  --build

logs:
	docker-compose logs -f

down:
	docker-compose -f docker-compose.dev.yml down

test:
	go test -v -cover ./...

clean: down
	@echo "=============cleaning up============="
	rm -rf web/build
	rm -rf web/node_modules
	rm -rf worker/worker
	docker system prune -f
	docker volume prune -f

init:
	cd infra/base && terraform init

up-instances: init
	@echo "=============instances spinning up============="
	cd infra/base && terraform apply -auto-approve -var "public_key=$$HOME/.ssh/id_rsa.pub" \
  														 -var "private_key=$$HOME/.ssh/id_rsa" \
  														 -var "ssh_fingerprint=$(SSH_FINGERPRINT)"

upload-inventory:
	cd infra/base && master=$$(terraform output master_ipv4_address) && scp inventory.txt root@$$master:/etc/ansible/inventory.txt

ansible-exec: upload-inventory
	cd infra/base && master=$$(terraform output master_ipv4_address) && ssh -t root@$$master 'cd /etc/ansible && ansible-playbook -i inventory.txt docker-playbook.yml'

destroy:
	cd infra/base && terraform destroy -auto-approve -var "public_key=$$HOME/.ssh/id_rsa.pub" \
  														   -var "private_key=$$HOME/.ssh/id_rsa" \
  														   -var "ssh_fingerprint=$(SSH_FINGERPRINT)"

finger:
	@echo this is my fingerprint $(SSH_FINGERPRINT)

output:
	cd infra/base && terraform output

	