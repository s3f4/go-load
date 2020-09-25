#!/bin/bash

function removeNodes() {
    for i in {1..3}; do
        docker-machine rm node$i -y
    done
}

function createNodes() {
    for i in {1..3}; do
        docker-machine create --driver virtualbox --virtualbox-disk-size "40000" --virtualbox-memory 2048 node$i
        #docker-machine create -d virtualbox node$i
    done
}

function joinSwarm() {
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
    docker node update --label-add role=master node1
    docker node update --label-add role=data node2
    docker node update --label-add role=worker node3
}


function moveFiles() {
    docker-machine ssh node1 "rm -rf /tmp/*"
    docker-machine ssh node1 "sudo rm -rf /app && sudo mkdir /app"
    docker-machine scp -r ./apigateway node1:/tmp/apigateway
    docker-machine scp -r ./eventhandler node1:/tmp/eventhandler
    rm -rf ./web/node_modules
    docker-machine scp -r ./web node1:/tmp/web
    docker-machine scp -r ./worker node1:/tmp/worker
    docker-machine scp -r ./swarm-dev.yml node1:/tmp/swarm-dev.yml
    
    docker-machine ssh node1 "
    sudo mv /tmp/apigateway /app/apigateway && \
    sudo mv /tmp/eventhandler /app/eventhandler && \
    sudo mv /tmp/web /app/web && \
    sudo mv /tmp/worker /app/worker &&
    sudo mv /tmp/swarm-dev.yml /app/swarm-dev.yml && exit"
}

function installCompose() {
    docker-machine ssh node1 sudo curl -L "https://github.com/docker/compose/releases/download/1.27.4/docker-compose-Linux-x86_64" -o /usr/local/bin/docker-compose
    docker-machine ssh node1 sudo chmod +x /usr/local/bin/docker-compose
    docker-machine ssh node1 sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
}

function deploySwarm() {
    docker-machine ssh node1 "MACHINE_HOST=$(docker-machine ip node1) && docker-compose -f /app/swarm-dev.yml up -d"
    docker-machine ssh node1 "MACHINE_HOST=$(docker-machine ip node1) && docker-compose -f /app/swarm-dev.yml down"
    docker-machine ssh node1 "MACHINE_HOST=$(docker-machine ip node1) &&docker stack deploy -c /app/swarm-dev.yml go-load"
}



while getopts ":h :r :c :j :m :i" opt; do
    case ${opt} in
        r )
            echo "Nodes are being removed..."
            removeNodes
            exit 0
        ;;
        c )
            echo "Nodes are being created..."
            exit 0
        ;;
        j )
            echo "Nodes are being joined to swarm..."
            joinSwarm
        ;;
        m )
            echo "Files are being moved..."
            moveFiles
        ;;
        i )
            echo "Docker compose are being installed..."
            installCompose
            echo ${OPTARG}
            exit 0
        ;;
        d )
            echo "Docker stack is being deployed..."
            deploySwarm
            echo ${OPTARG}
            exit 0
        ;;
        \? )
            echo "Invalid Option: -$OPTARG" 1>&2
            exit 1
        ;;
    esac
done
exit 0