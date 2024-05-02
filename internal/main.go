package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LeandroVCastro/applying-manager-api/internal/database"
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	"github.com/LeandroVCastro/applying-manager-api/internal/routes"

	"github.com/joho/godotenv"
)

func main() {
	errorGoDotEnv := godotenv.Load()
	if errorGoDotEnv != nil {
		log.Fatal("Error loading .env file")
	}

	database.StartConnection()

	var user entity.User
	database.Connection.Find(&user)
	fmt.Println("users: ", user)

	// fmt.Println("Database connection: ", db)
	routesRegister := routes.RunApi()
	fmt.Println("Applying manager API is running on :8081 port")
	log.Fatal(http.ListenAndServe(":8081", routesRegister))
}
