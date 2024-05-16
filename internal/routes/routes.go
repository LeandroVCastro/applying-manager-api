package routes

import (
	application "github.com/LeandroVCastro/applying-manager-api/internal/application/company"
	"github.com/LeandroVCastro/applying-manager-api/internal/routes/middleware"
	"github.com/gorilla/mux"
)

func RunApi() *mux.Router {
	muxRouter := mux.NewRouter()

	// Company routes
	companyRoutes := muxRouter.PathPrefix("/company").Subrouter()
	companyRoutes.HandleFunc("", middleware.DbTransactions(application.SaveCompany)).Methods("POST")
	companyRoutes.HandleFunc("", middleware.DbTransactions(application.ListCompanies)).Methods("GET")

	return muxRouter
}
