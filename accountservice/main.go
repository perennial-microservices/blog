package main

import (
	"fmt"
	"github.com/perennial-microservices/blog/accountservice/service"
	"github.com/perennial-microservices/blog/accountservice/dbclient"
)

var serviceName = "Account Service"

func main() {
	fmt.Printf("Starting %v\n", serviceName)
	initializeBoltClient()
	service.StartWebServer("6767")
}

func initializeBoltClient() {
	//service.DBClient = &dbclient.BoltClient{}
	//service.DBClient.OpenBoltDb()
	service.DBClient = &dbclient.RedisClient{}
	service.DBClient.StartRedis()
	service.DBClient.Seed()
}