package application_platform

import (
	"net/http"
	"strconv"

	platform_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/platform"
	"github.com/LeandroVCastro/applying-manager-api/internal/routes/route_handlers"
	"github.com/gorilla/mux"
)

func DeletePlatform(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	platformId, err := strconv.Atoi(vars["id"])
	if err != nil {
		route_handlers.ErrorResponse(response, request, "invalid ID", 400)
		return
	}
	deletePlatform := platform_domain.DeletePlatformFactory()
	errStatus, errDeletePlatform := deletePlatform.Handle(uint(platformId))
	if errDeletePlatform != nil {
		route_handlers.ErrorResponse(response, request, errDeletePlatform.Error(), errStatus)
		return
	}
	route_handlers.SuccessResponse(response, nil, http.StatusOK)
}
