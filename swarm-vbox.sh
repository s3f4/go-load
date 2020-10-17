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
    joinSwarm
    moveFiles
    registry
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

function registry() {
    ## create cert files
    echo "cert files are being created"
    openssl req -newkey rsa:4096 -nodes -sha256 \
    -keyout registry.key -x509 -days 365 \
    -out registry.crt -subj '/C=TR/ST=TR/L=Malatya/O=registry/CN=registry.dev'
    
    echo "files are being moved to nodes"
    
    
    for i in {1..3}; do
        if [ "$i" -eq 1 ]; then
            docker-machine scp -r ./registry.key node$i:/tmp/registry.key
            docker-machine ssh node$i sudo mv /tmp/registry.key /home/docker
        fi
        docker-machine scp -r ./registry.crt node$i:/tmp/registry.crt
        docker-machine ssh node$i sudo mv /tmp/registry.crt /home/docker
        docker-machine ssh node$i "sudo -- sh -c 'echo $(docker-machine ip node1) registry.dev >> /etc/hosts'"
        docker-machine ssh node$i sudo mkdir -p /etc/docker/certs.d/registry.dev:5000/
        docker-machine ssh node$i sudo cp /home/docker/registry.crt /etc/docker/certs.d/registry.dev:5000/ca.crt
    done
    
    docker-machine ssh node1 docker service create -d --name registry --publish=5000:5000 \
    --constraint=node.role==manager \
    --mount=type=bind,src=/home/docker,dst=/certs \
    -e REGISTRY_HTTP_ADDR=0.0.0.0:5000 \
    -e REGISTRY_HTTP_TLS_CERTIFICATE=/certs/registry.crt \
    -e REGISTRY_HTTP_TLS_KEY=/certs/registry.key \
    registry:latest

    docker-machine ssh node1 "ssh-keygen -t rsa -b 4096 -N '' -C "sefa@dehaa.com" -f ~/.ssh/id_rsa_for_master"
    
    docker-machine ssh node1 docker build -t registry.dev:5000/apigateway /app/apigateway -f /app/apigateway/Dockerfile.dev
    docker-machine ssh node1 docker push registry.dev:5000/apigateway
    docker-machine ssh node1 docker build -t registry.dev:5000/worker /app/worker -f /app/worker/Dockerfile.dev
    docker-machine ssh node1 docker push registry.dev:5000/worker
    docker-machine ssh node1 docker build -t registry.dev:5000/web /app/web -f /app/web/Dockerfile.dev
    docker-machine ssh node1 docker push registry.dev:5000/web
    docker-machine ssh node1 docker build -t registry.dev:5000/eventhandler /app/eventhandler -f /app/eventhandler/Dockerfile.dev
    docker-machine ssh node1 docker push registry.dev:5000/eventhandler
    
    
}


function moveFiles() {
    docker-machine ssh node1 "rm -rf /tmp/*"
    docker-machine ssh node1 "sudo rm -rf /app && sudo mkdir /app"
    docker-machine scp -r ./apigateway node1:/tmp/apigateway
    docker-machine scp -r ./eventhandler node1:/tmp/eventhandler
    rm -rf ./web/node_modules
    echo "REACT_APP_API_BASE_URL=$(docker-machine ip node1):3001" > ./web/.env
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
    docker-machine ssh node1 "docker stack rm go-load"
    docker-machine ssh node1 "docker-compose -f /app/swarm-dev.yml up --build -d"
    docker-machine ssh node1 "docker-compose -f /app/swarm-dev.yml down"
    docker-machine ssh node1 "docker stack deploy -c /app/swarm-dev.yml go-load"
}



while getopts ":h :r :c :j :m :i :f :s" opt; do
    case ${opt} in
        f)
            echo "Files are being moved and stack will be restarted"
            moveFiles
            deploySwarm
        ;;
        s)
            echo "Installing is getting started from zero"
            removeNodes
            createNodes
            joinSwarm
            installCompose
            moveFiles
            deploySwarm
            exit 0
        ;;
        r )
            echo "Nodes are being removed..."
            removeNodes
            exit 0
        ;;
        c )
            echo "Nodes are being created..."
            createNodes
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
echo $(docker-machine ip node1)
exit 0