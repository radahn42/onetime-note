package validator

import (
	"net/url"
	"strings"

	"github.com/go-playground/validator/v10"
)

func validateRelativeURL(fl validator.FieldLevel) bool {
	s := fl.Field().String()

	if !strings.HasPrefix(s, "/") {
		return false
	}

	u, err := url.ParseRequestURI(s)
	return err == nil && u.Scheme == "" && u.Host == ""
}
