package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/cuffymate1/pos-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func ConnDb() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	p, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Invalid port number: %s", port)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, p, user, password, dbname)
	// New logger for detailed SQL logging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&models.Product{}, &models.Category{}, &models.Users{}, &models.Order{}, &models.OrderItem{}, &models.Topping{}, &models.OrderItemTopping{}, &models.Payment{})

	fmt.Println("connect to database completed!")
}

func GetDb() *gorm.DB {
	return db
}
