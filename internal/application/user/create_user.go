package application

import (
	"fmt"
	"net/http"

	user_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/user"
	"github.com/LeandroVCastro/applying-manager-api/internal/routes/route_handlers"
)

type RequestBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func CreateUser(response http.ResponseWriter, request *http.Request) {
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

	createUser := user_domain.CreateUserFactory()
	createdUser, errCreateUser := createUser.Handle(body.Email, body.Password)
	if errCreateUser != nil {
		route_handlers.ErrorResponse(response, errCreateUser.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("createdUser", createdUser)
	fmt.Fprintf(response, "User Created: ")
}
