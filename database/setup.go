package database

import (
	"fmt"
	"log"
	"os"

	"github.com/SzRoland13/todo-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DB_DSN")

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database connected!")

	err = DB.AutoMigrate(&models.User{}, &models.Todo{})
	if err != nil {
		log.Fatal("Failed to auto-migrate:", err)
	}
}
