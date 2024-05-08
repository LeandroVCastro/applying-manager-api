package routes

import (
	"fmt"
	"net/http"

	application "github.com/LeandroVCastro/applying-manager-api/internal/application/company"
	"github.com/gorilla/mux"
)

func RunApi() *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/hello", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "Olá, mundo!")
	}).Methods("GET")

	muxRouter.HandleFunc("/company", application.SaveCompany).Methods("POST")
	return muxRouter
}
