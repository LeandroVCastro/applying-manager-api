package applyment_application

import (
	"net/http"
	"time"

	applyment_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/applyment"
	"github.com/LeandroVCastro/applying-manager-api/internal/routes/route_handlers"
)

type RequestBody struct {
	ID          uint      `json:"id" validate:"omitempty,number"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"omitempty"`
	Link        string    `json:"link" validate:"omitempty,fqdn|http_url"`
	PlatformId  uint      `json:"platform_id" validate:"omitempty,number"`
	CompanyId   uint      `json:"company_id" validate:"omitempty,number"`
	AppliedAt   time.Time `json:"applied_at" validate:"omitempty,datetime"`
}

func SaveApplyment(response http.ResponseWriter, request *http.Request) {
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
	saveApplyment := applyment_domain.SaveApplymentFactory()
	applyment, errStatus, err := saveApplyment.Handle(
		body.ID,
		body.Title,
		&body.Description,
		&body.Link,
		&body.CompanyId,
		&body.PlatformId,
		&body.AppliedAt,
	)
	if err != nil {
		route_handlers.ErrorResponse(response, request, err.Error(), errStatus)
		return
	}
	route_handlers.SuccessResponse(response, applyment, http.StatusOK)
}
