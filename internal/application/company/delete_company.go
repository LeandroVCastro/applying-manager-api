package application_company

import (
	"net/http"
	"strconv"

	company_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/company"
	"github.com/LeandroVCastro/applying-manager-api/internal/routes/route_handlers"
	"github.com/gorilla/mux"
)

func DeleteCompany(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	companyId, err := strconv.Atoi(vars["id"])
	if err != nil {
		route_handlers.ErrorResponse(response, request, "invalid ID", 400)
		return
	}
	deleteCompany := company_domain.DeleteCompanyFactory()
	errStatus, errDeleteCompany := deleteCompany.Handle(uint(companyId))
	if errDeleteCompany != nil {
		route_handlers.ErrorResponse(response, request, errDeleteCompany.Error(), errStatus)
		return
	}
	route_handlers.SuccessResponse(response, nil, http.StatusOK)
}
