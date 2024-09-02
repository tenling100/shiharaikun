package infrastructure

import (
	"context"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/tenling100/shiharaikun/api"
	"github.com/tenling100/shiharaikun/internal/config"
	interface_grpc "github.com/tenling100/shiharaikun/internal/interface/grpc"
)

type server struct {
	server *grpc.Server
	env    *config.Env
}

func NewGRPCServer(
	ctx context.Context,
	env *config.Env,
) *server {
	grpcServer := grpc.NewServer()

	return &server{
		server: grpcServer,
		env:    env,
	}
}

func (s *server) Run(
	invoiceServer *interface_grpc.InvoiceServer,
) error {
	log.Default().Println("Starting GRPC server on port " + s.env.GRPCPort)
	pb.RegisterInvoiceServiceServer(s.server, invoiceServer)
	reflection.Register(s.server)
	lis, err := net.Listen("tcp", ":"+s.env.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	if err := s.server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}

	// start http server
	log.Default().Println("Starting HTTP server on port " + s.env.HTTPPort)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	err = pb.RegisterInvoiceServiceHandlerServer(ctx, mux, invoiceServer)
	if err != nil {
		log.Fatalf("failed to register server: %v", err)
		return err
	}
	if err := http.ListenAndServe(":"+s.env.HTTPPort, mux); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}

	return nil
}

func (s *server) Shutdown() {
	log.Default().Println("Shutting down GRPC server")
	s.server.GracefulStop()
}
