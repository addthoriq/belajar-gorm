package models

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := "localhost"
	user := "postgres"
	password := "123"
	dbname := "belajar_gorm_db"
	port := "5432"
	sslmode := "disable"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect Database: %v", err)
	}

	db.AutoMigrate(&User{})

	// Setup Pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed get Pool: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)                  // 10 pool idle
	sqlDB.SetMaxOpenConns(100)                 // 100 pool max open
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)  // max idle time
	sqlDB.SetConnMaxLifetime(60 * time.Minute) // max life time

	DB = db

	log.Println("Connected Database")
}
