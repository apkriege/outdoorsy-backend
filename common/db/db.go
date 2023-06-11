package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Init initializes the database connection
func Init() error {
	dsn := "host=localhost user=root password=root dbname=testingwithrentals port=5434 sslmode=disable"

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	return nil
}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	fmt.Println("Get DB: ", db)
	return db
}

// CloseDBConnection closes the database connection
func CloseDBConnection() error {
	db, err := GetDB().DB()
	if err != nil {
		return fmt.Errorf("failed to close the database connection: %w", err)
	}

	return db.Close()
}
