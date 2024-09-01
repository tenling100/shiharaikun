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

const (
	maxIdleConns = 10
	maxOpenConns = 100
)

func InitializeDB(env *config.Env) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.DBUser,
		env.DBPass,
		env.DBHost,
		env.DBPort,
		env.DBName,
	)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return nil, err
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get DB instance: %v", err)
		return nil, err
	}

	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)

	return DB, nil
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&domain.ClientBankAccount{},
		&domain.Company{},
		&domain.InvoiceData{},
		&domain.User{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
		return err
	}
	return nil
}

func CloseDB(db *gorm.DB) error {
	if db == nil {
		log.Fatalf("failed to close database: DB is nil")
		return errors.New("DB is nil")
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get DB instance: %v", err)
		return err
	}
	sqlDB.Close()
	return nil
}
