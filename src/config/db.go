package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB initializes the global DB connection with connection pooling
func ConnectDB() (*gorm.DB, error) {
	dsn := confMysql()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
		return nil, err
	}

	// Get the underlying sql.DB instance to configure connection pooling
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Could not get sql.DB from Gorm DB: %v", err)
		return nil, err
	}

	// Configure the connection pool
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}

func confMysql() string {
	username, password, hostname, dbname := os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME")
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, hostname, dbname)
}
