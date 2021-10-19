package helper

import (
	validator "github.com/go-playground/validator/v10"
)

// CustomValidator : custom json body binding validator
type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

