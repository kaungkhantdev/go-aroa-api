package utils

import (
    "github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// InitValidator initializes the validator instance
func InitValidator() {
    validate = validator.New()
}

// ValidateStruct is a reusable function for struct validation
func ValidateStruct(input interface{}) error {
    return validate.Struct(input)
}
