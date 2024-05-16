package application

import (
	"net/http"

	company_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/company"
	"github.com/LeandroVCastro/applying-manager-api/internal/routes/route_handlers"
)

func ListCompanies(response http.ResponseWriter, request *http.Request) {
	listCompanies := company_domain.ListCompaniesFactory()
	listedCompanies, errStatus, err := listCompanies.Handle()
	if err != nil {
		route_handlers.ErrorResponse(response, request, err.Error(), errStatus)
		return
	}
	route_handlers.SuccessResponse(response, listedCompanies, http.StatusOK)
}
