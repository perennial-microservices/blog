package service

import (
	"fmt"
	"github.com/perennial-microservices/blog/accountservice/db"
)

var serviceName = "Account Service"

func Start() {
	fmt.Println("Starting ", serviceName)
	bootstrap()
}

func bootstrap() {
	DBClient = &db.AccountsClient{}
	DBClient.OpenBoltDb(db.BucketName)
	DBClient.Seed()
}
