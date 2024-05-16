package application_company

import (
	"net/http"
	"strconv"

	company_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/company"
	"github.com/LeandroVCastro/applying-manager-api/internal/routes/route_handlers"
	"github.com/gorilla/mux"
)

func GetCompany(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	companyId, err := strconv.Atoi(vars["id"])
	if err != nil {
		route_handlers.ErrorResponse(response, request, "invalid ID", 400)
		return
	}
	getCompany := company_domain.GetCompanyFactory()
	company, errStatus, errGetCompany := getCompany.Handle(uint(companyId))
	if errGetCompany != nil {
		route_handlers.ErrorResponse(response, request, errGetCompany.Error(), errStatus)
		return
	}
	route_handlers.SuccessResponse(response, company, http.StatusOK)
}
