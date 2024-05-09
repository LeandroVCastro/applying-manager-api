package routes

import (
	application "github.com/LeandroVCastro/applying-manager-api/internal/application/company"
	"github.com/LeandroVCastro/applying-manager-api/internal/routes/middleware"
	"github.com/gorilla/mux"
)

func RunApi() *mux.Router {
	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/company", middleware.DbTransactions(application.SaveCompany)).Methods("POST")
	return muxRouter
}
