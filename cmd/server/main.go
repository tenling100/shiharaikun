package main

import (
	"context"
	"log"

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
	grpcServer := infrastructure.NewGRPCServer(ctx, env)

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
	grpcServer.Run(invoiceServer)
	if err != nil {
		log.Fatalf("Failed to run gRPC server: %v", err)
	}

	// gracefully shutdown the server
	grpcServer.Shutdown()

}
