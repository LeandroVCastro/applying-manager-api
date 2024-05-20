package routes

import (
	application_company "github.com/LeandroVCastro/applying-manager-api/internal/application/company"
	application_platform "github.com/LeandroVCastro/applying-manager-api/internal/application/platform"
	"github.com/LeandroVCastro/applying-manager-api/internal/routes/middleware"
	"github.com/gorilla/mux"
)

func RunApi() *mux.Router {
	muxRouter := mux.NewRouter()

	// Company routes
	companyRoutes := muxRouter.PathPrefix("/company").Subrouter()
	companyRoutes.HandleFunc("", middleware.DbTransactions(application_company.SaveCompany)).Methods("POST")
	companyRoutes.HandleFunc("", middleware.DbTransactions(application_company.ListCompanies)).Methods("GET")
	companyRoutes.HandleFunc("/{id}", middleware.DbTransactions(application_company.GetCompany)).Methods("GET")
	companyRoutes.HandleFunc("/{id}", middleware.DbTransactions(application_company.DeleteCompany)).Methods("DELETE")

	// Platform routes
	platformRoutes := muxRouter.PathPrefix("/platform").Subrouter()
	platformRoutes.HandleFunc("", middleware.DbTransactions(application_platform.SavePlatform)).Methods("POST")
	platformRoutes.HandleFunc("", middleware.DbTransactions(application_platform.ListPlatforms)).Methods("GET")
	// companyRoutes.HandleFunc("/{id}", middleware.DbTransactions(application_company.GetCompany)).Methods("GET")
	// companyRoutes.HandleFunc("/{id}", middleware.DbTransactions(application_company.DeleteCompany)).Methods("DELETE")

	return muxRouter
}
