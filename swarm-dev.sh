#!/bin/bash

for i in {1..3}; do
    docker-machine rm node$i -y
done

for i in {1..3}; do
    docker-machine create -d virtualbox node$i
done

masterIP=$(docker-machine ip node1)
dataIP=$(docker-machine ip node2)

eval $(docker-machine env node1)
docker swarm init --advertise-addr $masterIP

workerToken=$(docker swarm join-token worker -q)

for i in {2..3}; do
    eval $(docker-machine env node$i)
    docker swarm join --token $workerToken $masterIP:2377
done

eval $(docker-machine env node1)
docker node update --label-add role=manager node1
docker node update --label-add role=data node2
docker node update --label-add role=worker node2

docker-machine ssh node1 sudo mkdir /app
docker-machine scp -r ./apigateway node1:/tmp/apigateway
docker-machine scp -r ./eventhandler node1:/tmp/eventhandler
docker-machine scp -r ./web node1:/tmp/web
docker-machine scp -r ./worker node1:/tmp/worker

docker-machine ssh node1 sudo mv /tmp/apigateway /app/apigateway
docker-machine ssh node1 sudo mv /tmp/eventhandler /app/eventhandler
docker-machine ssh node1 sudo mv /tmp/web /app/web
docker-machine ssh node1 sudo mv /tmp/worker /app/worker
