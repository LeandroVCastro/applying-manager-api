package middleware

import (
	"fmt"
	"net/http"

	"github.com/LeandroVCastro/applying-manager-api/internal/configs"
	"github.com/LeandroVCastro/applying-manager-api/internal/routes/route_handlers"
)

func DbTransactions(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		transactions := configs.Connection.Begin()
		next(response, request)
		if request.Header.Get(route_handlers.HeaderSuccessRequest) == "false" {
			fmt.Println("transactions rollback: ", request.URL)
			transactions.Rollback()
		} else {
			fmt.Println("transactions commit: ", request.URL)
			transactions.Commit()
		}
	}
}
