package service

import (
	"fmt"

	"github.com/gin-gonic/gin"

	pb "protoc-gen-gin-http/example/api/greeter"
)

type GreeterService struct{}

var _ pb.GreeterHTTPServer = (*GreeterService)(nil)

func (s *GreeterService) SayHello(c *gin.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello, %s! You are %d years old.", req.Name, req.Age),
	}, nil
}

func (s *GreeterService) CreateUser(c *gin.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{
		Id:    1,
		Name:  req.Name,
		Email: req.Email,
	}, nil
}

func (s *GreeterService) GetUser(c *gin.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{
		Id:    req.Id,
		Name:  "test-user",
		Email: "test@example.com",
		Age:   25,
	}, nil
}
