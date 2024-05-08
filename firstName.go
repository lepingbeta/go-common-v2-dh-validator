package dhvalidator

import (
	"github.com/go-playground/validator/v10"
)

func validateFirstName(firstName string) bool {
	// 定义正则表达式
	pattern := "^[\\p{L}][\\p{L}\\p{N}_-]{0,19}$"

	return validateRegex(firstName, pattern)
}

func IsValidFirstName(fl validator.FieldLevel) bool {
	firstName := fl.Field().String()
	return validateFirstName(firstName)
}
