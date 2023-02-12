#!/bin/bash
docker run --rm -p 8082:8090 --net=kvsnet --ip=10.10.0.2 --name main-instance kvsimg