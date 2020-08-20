SSH_FINGERPRINT=$(shell fp=`ssh-keygen -E md5 -lf ~/.ssh/id_rsa.pub | awk '{print $$2}'` && echo "$${fp/MD5:/}")

default:
	@echo "=============Building============="
	docker build -f worker/Dockerfile.dev -t worker .

up: default
	@echo "=============Compose Up============="
	docker-compose up -d

logs:
	docker-compose logs -f

down:
	docker-compose down

test:
	go test -v -cover ./...

clean: down
	@echo "=============cleaning up============="
	rm -f api
	docker system prune -f
	docker volume prune -f

up-instances:
	@echo "=============instances spinning up============="
	cd infra/base && terraform apply -auto-approve -var "public_key=$$HOME/.ssh/id_rsa.pub" \
  														 -var "private_key=$$HOME/.ssh/id_rsa" \
  														 -var "ssh_fingerprint=$(SSH_FINGERPRINT)"

destroy:
	cd infra/base && terraform destroy -auto-approve -var "public_key=$$HOME/.ssh/id_rsa.pub" \
  														   -var "private_key=$$HOME/.ssh/id_rsa" \
  														   -var "ssh_fingerprint=$(SSH_FINGERPRINT)"

finger:
	@echo this is my fingerprint $(SSH_FINGERPRINT)

output:
	cd infra/base && terraform output

	