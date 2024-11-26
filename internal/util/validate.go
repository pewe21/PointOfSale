package util

import (
	"github.com/go-playground/validator/v10"
)

func Validate(req interface{}) error {
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return err.(validator.ValidationErrors)

	}
	return nil
}
