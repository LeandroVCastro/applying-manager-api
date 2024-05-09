package route_handlers

import (
	"fmt"
	"net/http"
)

const HeaderSuccessRequest string = "Success-Request"

func ErrorResponse(response http.ResponseWriter, request *http.Request, message string, statusCode int) {
	response.WriteHeader(statusCode)
	request.Header.Add(HeaderSuccessRequest, "false")
	fmt.Fprintf(response, "Error: "+message)
}
