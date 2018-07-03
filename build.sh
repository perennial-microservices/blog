#!/bin/bash

cd accountservice;env GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags="-w -s" -o accountservice-linux-amd64;echo built `pwd`;cd ..;
cd healthcheck;env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o healthcheck-linux-amd64;echo built `pwd`;cd ..;

cp healthcheck/healthcheck-linux-amd64 accountservice/

docker build -t asahasrabuddh/account-service:2.0 accountservice/
docker push asahasrabuddh/account-service