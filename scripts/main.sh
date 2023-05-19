#!/bin/bash

# Clean up previous attempt
docker kill $(docker ps -a -q)
echo "y" | docker image prune -a
echo "y" | docker system prune -a

# Set up
docker build -t kvsimg .
docker network create --subnet=10.10.0.0/16 kvsnet

gnome-terminal --tab --title="Alice" -e "./scripts/alice.sh" --tab --title="Bob"  -e "./scripts/bob.sh"