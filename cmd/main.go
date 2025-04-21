package main

import (
	"log"
	"os"

	"github.com/SzRoland13/todo-backend/database"
	"github.com/SzRoland13/todo-backend/routes"
)

func main() {

	database.Connect()

	router := routes.SetupRouter()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Server running on http://localhost:" + port)
	router.Run(":" + port)

}
