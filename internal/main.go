package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LeandroVCastro/applying-manager-api/internal/configs"
	"github.com/LeandroVCastro/applying-manager-api/internal/routes"

	"github.com/joho/godotenv"
)

func main() {
	errorGoDotEnv := godotenv.Load()
	if errorGoDotEnv != nil {
		log.Fatal("Error loading .env file")
	}

	configs.StartConnection()
	routesRegister := routes.RunApi()
	fmt.Println("Applying manager API is running on :8081 port")
	log.Fatal(http.ListenAndServe(":8081", routesRegister))
}
