package main

import (
	"fmt"
	"github.com/perennial-microservices/blog/accountservice/service"
)

var serviceName = "Account Service"

func main() {
	fmt.Printf("Starting %v\n", serviceName)
	service.StartWebServer("6767")
}
