package application

import (
	"fmt"
	"net/http"

	userdomain "github.com/LeandroVCastro/applying-manager-api/internal/domain/user"
)

func GetUser(response http.ResponseWriter, request *http.Request) {
	createdUser := userdomain.GetUser()
	fmt.Fprintf(response, "Updated response: "+createdUser)
}
