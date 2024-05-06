package route_handlers

import (
	"fmt"
	"net/http"
)

func ErrorResponse(response http.ResponseWriter, message string, statusCode int) {
	response.WriteHeader(statusCode)
	fmt.Fprintf(response, "Error: "+message)
}
