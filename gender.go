package dhvalidator

import (
	"github.com/go-playground/validator/v10"
)


func validateGender(gender string) bool {
	// 定义正则表达式
	pattern := "^(?:male|female|other)$"

	return validateRegex(gender, pattern)
}

func IsValidGender(fl validator.FieldLevel) bool {
	gender := fl.Field().String()
	return validateGender(gender)
}
