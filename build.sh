#!/bin/bash
docker build -t registry.dev:5000/apigateway /root/app/apigateway -f /root/app/apigateway/Dockerfile.prod
docker push registry.dev:5000/apigateway
docker build -t registry.dev:5000/worker /root/app/worker -f /root/app/worker/Dockerfile.prod
docker push registry.dev:5000/worker
docker build -t registry.dev:5000/web /root/app/web -f /root/app/web/Dockerfile.prod
docker push registry.dev:5000/web
docker build -t registry.dev:5000/eventhandler /root/app/eventhandler -f /root/app/eventhandler/Dockerfile.prod
docker push registry.dev:5000/eventhandler