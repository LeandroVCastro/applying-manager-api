package platform_application

import (
	"net/http"

	platform_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/platform"
	"github.com/LeandroVCastro/applying-manager-api/internal/routes/route_handlers"
)

func ListPlatforms(response http.ResponseWriter, request *http.Request) {
	listPlatformsDomain := platform_domain.ListPlatformsFactory()
	listedPlatforms, errStatus, err := listPlatformsDomain.Handle()
	if err != nil {
		route_handlers.ErrorResponse(response, request, err.Error(), errStatus)
		return
	}
	route_handlers.SuccessResponse(response, listedPlatforms, http.StatusOK)
}
