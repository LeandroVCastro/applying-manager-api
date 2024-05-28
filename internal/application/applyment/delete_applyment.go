package applyment_application

import (
	"net/http"
	"strconv"

	applyment_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/applyment"
	"github.com/LeandroVCastro/applying-manager-api/internal/routes/route_handlers"
	"github.com/gorilla/mux"
)

func DeleteApplyment(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	applymentId, err := strconv.Atoi(vars["id"])
	if err != nil {
		route_handlers.ErrorResponse(response, request, "invalid ID", 400)
		return
	}
	deleteApplyment := applyment_domain.DeleteApplymentFactory()
	errStatus, err := deleteApplyment.Handle(uint(applymentId))
	if err != nil {
		route_handlers.ErrorResponse(response, request, err.Error(), errStatus)
		return
	}
	route_handlers.SuccessResponse(response, nil, http.StatusOK)
}
