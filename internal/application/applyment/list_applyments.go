package applyment_application

import (
	"net/http"

	applyment_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/applyment"
	"github.com/LeandroVCastro/applying-manager-api/internal/routes/route_handlers"
)

func ListApplyments(response http.ResponseWriter, request *http.Request) {
	listApplyments := applyment_domain.ListApplymentsFactory()
	applyments, errStatus, err := listApplyments.Handle()
	if err != nil {
		route_handlers.ErrorResponse(response, request, err.Error(), errStatus)
		return
	}
	route_handlers.SuccessResponse(response, applyments, http.StatusOK)
}
