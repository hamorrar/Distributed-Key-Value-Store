#!/bin/bash

# Clean up previous attempt
y | docker image prune -a
docker kill $(docker ps -q)
docker rm $(docker ps -a -q)
docker network rm kvsnet

# Set up
docker build -t kvsimg .
docker network create --subnet=10.10.0.0/16 kvsnet

# Run
docker run --rm -p 8082:8090 --net=kvsnet --ip=10.10.0.2 --name main-instance kvsimg
