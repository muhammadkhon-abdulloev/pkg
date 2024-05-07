package utils

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func Init() {
	validate = validator.New()
}

func ValidateStruct(a any) error {
	if validate == nil {
		Init()
	}

	if err := validate.Struct(a); err != nil {
		return err
	}

	return nil
}
