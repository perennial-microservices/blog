#!/bin/bash

cd accountservice;env GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o accountservice-linux-amd64;echo built `pwd`;cd ..;
cd healthcheck;env GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o healthcheck-linux-amd64;echo built `pwd`;cd ..;

cp healthcheck/healthcheck-linux-amd64 accountservice/

docker build -t asahasrabuddh/account-service:1.0 accountservice/
docker push asahasrabuddh/account-service
