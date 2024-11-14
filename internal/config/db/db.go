package db

import (
	"cmarin20/dnq-ecommerce/internal/app/products"
	"cmarin20/dnq-ecommerce/pkg/logger"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSLMode  string
}

func NewDbConfig() DbConfig {
	return DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
}

func NewDbConn(cfg DbConfig, logger *logger.Logger) *gorm.DB {
	fmt.Println(cfg)
	logger.Info("Opening a new database connection...")
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName, cfg.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		logger.Error("error to open database connection: %v", err)
		return nil
	}

	logger.Info("Migrating tables...")
	if err := db.AutoMigrate(&products.Product{}); err != nil {
		logger.Error("error during migration: %v", err)
		return nil
	}

	logger.Info("Database connection established.")
	return db
}

func CloseDbConn(db *gorm.DB, logger *logger.Logger) error {
	logger.Info("Closing the database connection...")
	dbSQL, err := db.DB()
	if err != nil {
		logger.Error("error to close database connection: %v", err)
		return err
	}
	dbSQL.Close()
	logger.Info("Database connection closed.")
	return nil
}
