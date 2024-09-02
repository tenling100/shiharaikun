package infrastructure

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/tenling100/shiharaikun/api"
	"github.com/tenling100/shiharaikun/internal/config"
	interface_grpc "github.com/tenling100/shiharaikun/internal/interface/grpc"
)

type server struct {
	grpcServer *grpc.Server
	httpServer *http.Server
	env        *config.Env
}

func NewServer(
	ctx context.Context,
	env *config.Env,
) *server {
	grpcServer := grpc.NewServer()
	httpServer := &http.Server{
		Addr:    ":" + env.HTTPPort,
		Handler: runtime.NewServeMux(),
	}
	return &server{
		grpcServer: grpcServer,
		httpServer: httpServer,
		env:        env,
	}
}

func (s *server) RunGRPC(
	invoiceServer *interface_grpc.InvoiceServer,
) error {
	log.Default().Println("Starting GRPC server on port " + s.env.GRPCPort)
	pb.RegisterInvoiceServiceServer(s.grpcServer, invoiceServer)
	reflection.Register(s.grpcServer)
	lis, err := net.Listen("tcp", ":"+s.env.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	if err := s.grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}
	return nil
}

func (s *server) RunHTTP(
	invoiceServer *interface_grpc.InvoiceServer,
) error {
	// start http server
	log.Default().Println("Starting HTTP server on port " + s.env.HTTPPort)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	err := pb.RegisterInvoiceServiceHandlerServer(ctx, s.httpServer.Handler.(*runtime.ServeMux), invoiceServer)
	if err != nil {
		log.Fatalf("failed to register server: %v", err)
		return err
	}
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running"))
	})

	if err := http.ListenAndServe(":"+s.env.HTTPPort, s.httpServer.Handler.(*runtime.ServeMux)); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}
	return nil
}

func (s *server) GRPCShutdown() {
	log.Default().Println("Shutting down GRPC server")
	s.grpcServer.GracefulStop()
}

func (s *server) HTTPShutdown() {
	log.Default().Println("Shutting down HTTP server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.httpServer.Shutdown(ctx)
}
