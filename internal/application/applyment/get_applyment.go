package applyment_application

import (
	"net/http"
	"strconv"

	applyment_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/applyment"
	"github.com/LeandroVCastro/applying-manager-api/internal/routes/route_handlers"
	"github.com/gorilla/mux"
)

func GetApplyment(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	applymentId, err := strconv.Atoi(vars["id"])
	if err != nil {
		route_handlers.ErrorResponse(response, request, "invalid ID", 400)
		return
	}
	getApplyment := applyment_domain.GetApplymentFactory()
	applyment, errStatus, errGetApplyment := getApplyment.Handle(uint(applymentId))
	if errGetApplyment != nil {
		route_handlers.ErrorResponse(response, request, errGetApplyment.Error(), errStatus)
		return
	}
	route_handlers.SuccessResponse(response, applyment, http.StatusOK)
}
