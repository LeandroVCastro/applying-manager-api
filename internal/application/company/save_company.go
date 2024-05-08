package application

import (
	"net/http"

	company_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/company"
	"github.com/LeandroVCastro/applying-manager-api/internal/routes/route_handlers"
)

type RequestBody struct {
	ID          uint   `json:"id" validate:"omitempty,number"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"omitempty"`
	Website     string `json:"website" validate:"omitempty,fqdn|http_url"`
	Linkedin    string `json:"linkedin" validate:"omitempty,fqdn|http_url"`
	Glassdoor   string `json:"glassdoor" validate:"omitempty,fqdn|http_url"`
	Instagram   string `json:"instagram" validate:"omitempty,fqdn|http_url"`
}

func SaveCompany(response http.ResponseWriter, request *http.Request) {
	body := RequestBody{}
	errBody := route_handlers.GetRequestBody(request, &body)
	if errBody != nil {
		route_handlers.ErrorResponse(response, errBody.Error(), http.StatusBadRequest)
		return
	}
	errValidateBody := route_handlers.ValidateBody(body)
	if errValidateBody != nil {
		route_handlers.ErrorResponse(response, errValidateBody.Error(), http.StatusBadRequest)
		return
	}
	saveCompany := company_domain.SaveCompanyFactory()
	savedCompany, err := saveCompany.Handle(
		body.ID,
		body.Name,
		&body.Description,
		&body.Website,
		&body.Linkedin,
		&body.Glassdoor,
		&body.Instagram,
	)
	if err != nil {
		route_handlers.ErrorResponse(response, err.Error(), http.StatusNotFound)
		return
	}
	route_handlers.SuccessResponse(response, savedCompany, http.StatusOK)
}