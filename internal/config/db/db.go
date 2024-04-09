package db

import (
	userModel "cmarin20/dnq-ecommerce/internal/user/model"
	"cmarin20/dnq-ecommerce/pkg/logger"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func NewDbConfig() DbConfig {
	return DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
	}
}

func NewDbConn(cfg DbConfig, logger *logger.Logger) *gorm.DB {
	logger.Info("Opening a new database connection...")
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		logger.Error("error to open database connection: %v", err)
		return nil
	}

	logger.Info("Migrating tables...")
	db.AutoMigrate(&userModel.User{})

	logger.Info("Database connection established.")
	return db
}

func CloseDbConn(db *gorm.DB, logger *logger.Logger) error {
	logger.Info("Closing the database connection...")
	dbSQL, err := db.DB()
	if err != nil {
		logger.Error("error to close database connection: %v", err)
	}
	dbSQL.Close()
	return err
}
