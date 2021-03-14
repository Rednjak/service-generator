#!/bin/bash

function stop_docker_containers {
  docker-compose -f "$(pwd)/deployments/dev/dependencies/docker-compose.yml" stop
}

go clean -i $(pwd)/cmd/server
go install $(pwd)/cmd/server

echo -e "\033[0;36mStarting backend server...\033[0m"

# start DB etc.
docker-compose -f "$(pwd)/deployments/dev/dependencies/docker-compose.yml" up -d

# Stop the containers when we provide "stop" as the second argument
if [ ! -z $1 ] && [ $1=='stop' ]; then
    trap stop_docker_containers EXIT
fi

server
