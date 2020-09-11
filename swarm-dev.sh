#!/bin/bash

for i in {1..4}; do
    docker-machine create -d virtualbox node$i
done

masterIP=$(docker-machine ip node1)
dataIP=$(docker-machine ip node2)

eval $(docker-machine env node1)
docker swarm init --advertise-addr $masterIP

workerToken=$(docker swarm join-token worker -q)

for i in {2..4}; do
    eval $(docker-machine env node$i)
    docker swarm join --token $workerToken $masterIP:2377
done
