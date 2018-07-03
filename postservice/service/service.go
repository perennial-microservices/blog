package service

import (
	"fmt"
	"github.com/perennial-microservices/blog/postservice/db"
)

var serviceName = "Post Service"

func Start() {
	fmt.Println("Starting ", serviceName)
	bootstrap()
}

func bootstrap() {
	DBClient = &db.PostsClient{}
	DBClient.OpenBoltDb(db.BucketName)
	DBClient.Seed()
}
