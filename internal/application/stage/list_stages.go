package stage_application

import (
	"net/http"

	stage_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/stage"
	"github.com/LeandroVCastro/applying-manager-api/internal/routes/route_handlers"
)

func ListStages(response http.ResponseWriter, request *http.Request) {
	listStages := stage_domain.ListStagesFactory()
	stages, errStatus, err := listStages.Handle()
	if err != nil {
		route_handlers.ErrorResponse(response, request, err.Error(), errStatus)
		return
	}
	route_handlers.SuccessResponse(response, stages, http.StatusOK)
}
