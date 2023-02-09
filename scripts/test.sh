#!/bin/bash

# Testing for PUT
curl --request PUT --header "Content-Type: application/json" --write-out "\n%{http_code}\n" --data '{"value": 123}' http://10.10.0.2:8090/kvs/key1