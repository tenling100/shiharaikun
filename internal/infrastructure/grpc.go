package infrastructure

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/tenling100/shiharaikun/api"
	"github.com/tenling100/shiharaikun/internal/config"
)

type server struct {
	server *grpc.Server
	env    config.Env
}

func NewGRPCServer(
	ctx context.Context,
	env config.Env,
) *server {
	grpcServer := grpc.NewServer()

	return &server{
		server: grpcServer,
		env:    env,
	}
}

func (s *server) Run(
	invoiceServer pb.InvoiceServiceServer,
	userServer pb.UserServiceServer,
) error {
	log.Default().Println("Starting GRPC server on port " + s.env.GRPCPort)

	pb.RegisterInvoiceServiceServer(s.server, invoiceServer)
	pb.RegisterUserServiceServer(s.server, userServer)

	lis, err := net.Listen("tcp", ":"+s.env.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	if err := s.server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}
	return nil
}
