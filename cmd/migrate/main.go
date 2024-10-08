package main

import (
	"log"

	"github.com/tenling100/shiharaikun/internal/config"
	"github.com/tenling100/shiharaikun/internal/infrastructure"
)

func main() {
	// Initialize GORM and connect to the database
	env, err := config.SetupEnv()
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}
	db, err := infrastructure.InitializeDB(env)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer infrastructure.CloseDB(db)

	// Auto-migrate the schema for all related models
	err = infrastructure.AutoMigrate(db)
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	log.Println("Database schema created/updated successfully!")
}
