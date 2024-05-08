package route_handlers

import (
	"encoding/json"
	"net/http"
)

func SuccessResponse(response http.ResponseWriter, content interface{}, statusCode int) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	json.NewEncoder(response).Encode(content)
}
