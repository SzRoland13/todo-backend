package main

import (
	"log"
	"os"

	"github.com/SzRoland13/todo-backend/database"
	"github.com/SzRoland13/todo-backend/routes"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env vars")
	}

	database.Connect()

	router := routes.SetupRouter()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Server running on http://localhost:" + port)
	router.Run(":" + port)

}
