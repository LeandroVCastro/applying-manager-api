package platform_application

import (
	"net/http"

	platform_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/platform"
	"github.com/LeandroVCastro/applying-manager-api/internal/routes/route_handlers"
)

type RequestBody struct {
	ID      uint   `json:"id" validate:"omitempty,number"`
	Name    string `json:"name" validate:"required"`
	Website string `json:"website" validate:"omitempty,fqdn|http_url"`
}

func SavePlatform(response http.ResponseWriter, request *http.Request) {
	body := RequestBody{}
	errBody := route_handlers.GetRequestBody(request, &body)
	if errBody != nil {
		route_handlers.ErrorResponse(response, request, errBody.Error(), http.StatusBadRequest)
		return
	}
	errValidateBody := route_handlers.ValidateBody(body)
	if errValidateBody != nil {
		route_handlers.ErrorResponse(response, request, errValidateBody.Error(), http.StatusBadRequest)
		return
	}
	savePlatform := platform_domain.SavePlatformFactory()
	savedPlatform, errStatus, err := savePlatform.Handle(
		body.ID,
		body.Name,
		&body.Website,
	)
	if err != nil {
		route_handlers.ErrorResponse(response, request, err.Error(), errStatus)
		return
	}
	route_handlers.SuccessResponse(response, savedPlatform, http.StatusOK)
}
