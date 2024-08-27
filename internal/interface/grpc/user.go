package grpc

import (
	"context"

	pb "github.com/tenling100/shiharaikun/api"
)

type userServer struct {
}

func NewUserServer() *userServer {
	return &userServer{}
}

func (s *userServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.UserResponse, error) {
	return nil, nil
}

func (s *userServer) GetUserById(ctx context.Context, in *pb.GetUserIdRequest) (*pb.UserResponse, error) {
	return nil, nil
}

func (s *userServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}

func (s *userServer) Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return nil, nil
}
