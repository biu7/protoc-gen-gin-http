package main

import (
	"github.com/gin-gonic/gin"

	pb "protoc-gen-gin-http/example/api/greeter"
	"protoc-gen-gin-http/example/service"
)

func main() {
	r := gin.Default()

	group := r.Group("")
	pb.RegisterGreeterHTTPServer(group, &service.GreeterService{})

	r.Run(":8080")
}
