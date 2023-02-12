#!/bin/bash
docker run --rm -p 8083:8090 --net=kvsnet --ip=10.10.0.3 -e FORWARDING_ADDRESS=10.10.0.2:8090 --name forwarding-instance1 kvsimg