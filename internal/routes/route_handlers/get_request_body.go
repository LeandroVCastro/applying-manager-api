package route_handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func GetRequestBody(request *http.Request, body interface{}) error {
	err := json.NewDecoder(request.Body).Decode(body)
	if err == nil {
		return nil
	}
	splited := strings.Split(err.Error(), "RequestBody.")
	if len(splited) > 1 {
		return errors.New(strings.Replace(splited[1], " of ", " should be ", 1))
	}
	return err
}
