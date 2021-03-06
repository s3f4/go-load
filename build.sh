#!/bin/bash

echo "cert files are being created"
openssl req -newkey rsa:4096 -nodes -sha256 \
-keyout /root/app/registry.key -x509 -days 365 \
-out /root/app/registry.crt -subj '/C=TR/ST=TR/L=LLL/O=registry/CN=registry.dev'

export ANSIBLE_HOST_KEY_CHECKING=False && ansible-playbook -i /etc/ansible/inventory.txt /etc/ansible/cert.yml

docker service create -d --name registry --publish=5000:5000 \
--constraint=node.role==manager \
--mount=type=bind,src=/root/app,dst=/certs \
-e REGISTRY_HTTP_ADDR=0.0.0.0:5000 \
-e REGISTRY_HTTP_TLS_CERTIFICATE=/certs/registry.crt \
-e REGISTRY_HTTP_TLS_KEY=/certs/registry.key \
registry:latest


docker build -t registry.dev:5000/apigateway /root/app/apigateway -f /root/app/apigateway/Dockerfile
docker push registry.dev:5000/apigateway
docker build -t registry.dev:5000/worker /root/app/worker -f /root/app/worker/Dockerfile
docker push registry.dev:5000/worker
docker build -t registry.dev:5000/web /root/app/web -f /root/app/web/Dockerfile
docker push registry.dev:5000/web
docker build -t registry.dev:5000/eventhandler /root/app/eventhandler -f /root/app/eventhandler/Dockerfile
docker push registry.dev:5000/eventhandler