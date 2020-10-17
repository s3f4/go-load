#!/bin/bash

echo "cert files are being created"
openssl req -newkey rsa:4096 -nodes -sha256 \
-keyout registry.key -x509 -days 365 \
-out registry.crt -subj '/C=TR/ST=TR/L=Malatya/O=registry/CN=registry.dev'


docker service create -d --name registry --publish=5000:5000 \
--constraint=node.role==manager \
--mount=type=bind,src=$(pwd),dst=/certs \
-e REGISTRY_HTTP_ADDR=0.0.0.0:5000 \
-e REGISTRY_HTTP_TLS_CERTIFICATE=/certs/registry.crt \
-e REGISTRY_HTTP_TLS_KEY=/certs/registry.key \
registry:latest

echo "You must add ' 127.0.0.1 registry.dev ' to /etc/hosts"
echo "Wait for starting registry service..."
sleep 2

docker build -t registry.dev:5000/apigateway ./apigateway -f ./apigateway/Dockerfile.dev
docker build -t registry.dev:5000/worker ./worker -f ./worker/Dockerfile.dev
docker build -t registry.dev:5000/web ./web -f ./web/Dockerfile.dev
docker build -t registry.dev:5000/eventhandler ./eventhandler -f ./eventhandler/Dockerfile.dev

docker push registry.dev:5000/apigateway
docker push registry.dev:5000/worker
docker push registry.dev:5000/web
docker push registry.dev:5000/eventhandler