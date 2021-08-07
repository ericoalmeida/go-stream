package services

import (
	"context"
	"fmt"

	"github.com/ericoalmeida/grpc/pb"
)

// This is an implementation of user.proto
// it must be registered on the server to work successfully

type UserService struct {
	pb.UnimplementedUserServiceServer
}

// It works like a constructor
func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, request *pb.User) (*pb.User, error) {
	fmt.Println(request.Name)

	return &pb.User{
		Id:    request.GetId(),
		Name:  request.GetName(),
		Email: request.GetEmail(),
	}, nil
}
