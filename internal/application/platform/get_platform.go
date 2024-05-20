package platform_application

import (
	"net/http"
	"strconv"

	platform_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/platform"
	"github.com/LeandroVCastro/applying-manager-api/internal/routes/route_handlers"
	"github.com/gorilla/mux"
)

func GetPlatform(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	platformId, err := strconv.Atoi(vars["id"])
	if err != nil {
		route_handlers.ErrorResponse(response, request, "invalid ID", 400)
		return
	}
	getPlatform := platform_domain.GetPlatformFactory()
	platform, errStatus, errGetPlatform := getPlatform.Handle(uint(platformId))
	if errGetPlatform != nil {
		route_handlers.ErrorResponse(response, request, errGetPlatform.Error(), errStatus)
		return
	}
	route_handlers.SuccessResponse(response, platform, http.StatusOK)
}
