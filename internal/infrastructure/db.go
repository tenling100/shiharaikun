package infrastructure

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/tenling100/shiharaikun/internal/config"
	"github.com/tenling100/shiharaikun/internal/domain"
)

var DB *gorm.DB

const (
	maxIdleConns = 10
	maxOpenConns = 100
)

func InitializeDB(env *config.Env) error {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.DBUser,
		env.DBPass,
		env.DBHost,
		env.DBPort,
		env.DBName,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return err
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get DB instance: %v", err)
		return err
	}

	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)

	return nil
}

func AutoMigrate() error {
	if DB == nil {
		log.Fatalf("failed to migrate database: DB is nil")
		return errors.New("DB is nil")
	}
	err := DB.AutoMigrate(
		&domain.ClientBankAccount{},
		&domain.Client{},
		&domain.Company{},
		&domain.InvoiceData{},
		&domain.Invoice{},
		&domain.User{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
		return err
	}
	return nil
}

func CloseDB() error {
	if DB == nil {
		log.Fatalf("failed to close database: DB is nil")
		return errors.New("DB is nil")
	}
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get DB instance: %v", err)
		return err
	}
	sqlDB.Close()
	return nil
}
