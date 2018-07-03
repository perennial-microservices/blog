package main

import (
	. "github.com/perennial-microservices/blog/service"
	accountService "github.com/perennial-microservices/blog/accountservice/service"
	postService "github.com/perennial-microservices/blog/postservice/service"
)

func main() {
	accountService.Start()
	postService.Start()

	AttachRoutes(accountService.AccountRoutes)
	AttachRoutes(postService.PostRoutes)

	StartWebServer("6767")
}
