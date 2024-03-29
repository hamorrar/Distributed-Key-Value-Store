#!/bin/bash

# Testing for PUT

# Basic, successful PUT to main
# curl --request PUT --header "Content-Type: application/json" --write-out "\n%{http_code}\n" --data '{"value": 123}' http://10.10.0.2:8090/kvs/key1

# Basic, successful PUT for forwarding
curl --request PUT --header "Content-Type: application/json" --write-out "\n%{http_code}\n" --data '{"value": 456}' http://10.10.0.2:8090/kvs/key2

# Basic, successful GET to main
# curl --request GET --header "Content-Type: application/json" --write-out "\n%{http_code}\n" http://10.10.0.2:8090/kvs/key1

# Basic, successful GET to forwarding
curl --request GET --header "Content-Type: application/json" --write-out "\n%{http_code}\n" http://10.10.0.3:8090/kvs/key2


# No Key Specified
# curl --request PUT --header "Content-Type: application/json" --write-out "\n%{http_code}\n" --data '{"value": 123}' http://10.10.0.2:8090/kvs/

# curl --request PUT --header "Content-Type: application/json" --write-out "\n%{http_code}\n" --data '{"value": 123}' http://10.10.0.2:8090/kvs