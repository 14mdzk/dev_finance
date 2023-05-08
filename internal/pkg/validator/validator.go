package validator

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Check(data interface{}) bool {
	err := validate.Struct(data)
	return err != nil
}
