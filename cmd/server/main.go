package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/tenling100/shiharaikun/internal/config"
	"github.com/tenling100/shiharaikun/internal/infrastructure"
	"github.com/tenling100/shiharaikun/internal/interface/grpc"
	"github.com/tenling100/shiharaikun/internal/repository"
	"github.com/tenling100/shiharaikun/internal/usecase"
)

func main() {
	ctx := context.Background()
	env, err := config.SetupEnv()
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}
	// Create a new gRPC server
	server := infrastructure.NewServer(ctx, env)

	// invoice server
	db, err := infrastructure.InitializeDB(env)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// close db connection
	defer infrastructure.CloseDB(db)

	invoiceDataRepository := repository.NewInvoiceDataRepository(db)
	companyRepository := repository.NewCompanyRepository(db)
	userRepository := repository.NewUserRepository(db)
	invoiceUseCase := usecase.NewInvoiceUsecaseImpl(invoiceDataRepository, companyRepository, userRepository, env)
	companyUseCase := usecase.NewCompanyUsecaseImpl(companyRepository)
	userUseCase := usecase.NewUserUsecaseImpl(userRepository)
	invoiceServer := grpc.NewInvoiceServer(invoiceUseCase, companyUseCase, userUseCase)

	// run grpc server
	go func() {
		err := server.RunGRPC(invoiceServer)
		if err != nil {
			log.Fatalf("Failed to run gRPC server: %v", err)
		}
		defer server.GRPCShutdown()
	}()
	// run http server
	go func() {
		err := server.RunHTTP(invoiceServer)
		if err != nil {
			log.Fatalf("Failed to run HTTP server: %v", err)
		}
	}()
	defer server.HTTPShutdown()

	// Handle graceful shutdown on SIGINT/SIGTERM
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	<-sigs
	log.Println("Shutting down gracefully...")
	// Perform any cleanup or shutdown tasks here
}
