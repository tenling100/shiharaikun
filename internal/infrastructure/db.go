package infrastructure

var DB *gorm.DB

const (
	maxIdleConns = 10
	maxOpenConns = 100
)

func InitializeDB(env config.Env) error {

	dns := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.DBHost,
		env.DBPort,
		env.DBUser,
		env.DBPass,
		env.DBName,
	)

	DB, err := gorm.Open(mysql.Open(dns), &gorm.Config{
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

	// Migrate the schema
	err = autoMigrate()
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
		return err
	}
	return nil
}

func autoMigrate() error {
	err := DB.AutoMigrate(
		&model.User{},
		&model.Invoice{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
		return err
	}
	return nil
}

func CloseDB() error {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get DB instance: %v", err)
		return err
	}
	sqlDB.Close()
	return nil
}
