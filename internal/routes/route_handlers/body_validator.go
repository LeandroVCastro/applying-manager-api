package route_handlers

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

func ValidateBody(body interface{}) error {
	errValidate := validate.Struct(body)
	if errValidate != nil {
		validationErrors := errValidate.(validator.ValidationErrors)
		for _, element := range validationErrors {
			return errors.New(strings.ToLower(element.Field()) + makeFriendly(element.Tag()))
		}
	}
	return nil
}

func makeFriendly(tag string) string {
	if tag == "required" {
		return " is required"
	}
	if tag == "email" {
		return " isn't a valid email"
	}
	if tag == "fqdn|http_url" {
		return " must to be an url"
	}
	return " " + tag
}
