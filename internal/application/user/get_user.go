package application

import (
	"fmt"
	"net/http"

	user_domain "github.com/LeandroVCastro/applying-manager-api/internal/domain/user"
)

func GetUser(response http.ResponseWriter, request *http.Request) {
	getUser := user_domain.GetUserFactory()
	user := getUser.Handle()
	fmt.Fprintf(response, "Updated response: "+user)
}
