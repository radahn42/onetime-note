package validator

import (
	"github.com/go-playground/validator/v10"
)

var v *validator.Validate

func init() {
	v = validator.New()
	v.RegisterValidation("relative_url", validateRelativeURL)
}

func Struct(s any) error {
	return v.Struct(s)
}
