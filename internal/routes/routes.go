package routes

import (
	company_application "github.com/LeandroVCastro/applying-manager-api/internal/application/company"
	platform_application "github.com/LeandroVCastro/applying-manager-api/internal/application/platform"
	"github.com/LeandroVCastro/applying-manager-api/internal/routes/middleware"
	"github.com/gorilla/mux"
)

func RunApi() *mux.Router {
	muxRouter := mux.NewRouter()

	// Company routes
	companyRoutes := muxRouter.PathPrefix("/company").Subrouter()
	companyRoutes.HandleFunc("", middleware.DbTransactions(company_application.SaveCompany)).Methods("POST")
	companyRoutes.HandleFunc("", middleware.DbTransactions(company_application.ListCompanies)).Methods("GET")
	companyRoutes.HandleFunc("/{id}", middleware.DbTransactions(company_application.GetCompany)).Methods("GET")
	companyRoutes.HandleFunc("/{id}", middleware.DbTransactions(company_application.DeleteCompany)).Methods("DELETE")

	// Platform routes
	platformRoutes := muxRouter.PathPrefix("/platform").Subrouter()
	platformRoutes.HandleFunc("", middleware.DbTransactions(platform_application.SavePlatform)).Methods("POST")
	platformRoutes.HandleFunc("", middleware.DbTransactions(platform_application.ListPlatforms)).Methods("GET")
	platformRoutes.HandleFunc("/{id}", middleware.DbTransactions(platform_application.GetPlatform)).Methods("GET")
	platformRoutes.HandleFunc("/{id}", middleware.DbTransactions(platform_application.DeletePlatform)).Methods("DELETE")

	return muxRouter
}
