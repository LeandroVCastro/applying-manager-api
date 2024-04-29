package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LeandroVCastro/applying-manager-api/internal/routes"
)

func main() {
	routesRegister := routes.RunApi()

	fmt.Println("Applying manager API is running on :8081 port")
	log.Fatal(http.ListenAndServe(":8081", routesRegister))
}
