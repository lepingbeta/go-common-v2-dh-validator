package dhvalidator

import (
	"github.com/go-playground/validator/v10"
)


func validateLastName(lastName string) bool {
	// 定义正则表达式
	pattern := "^[\\p{L}][\\p{L}\\p{N}_-]{0,19}$"

	return validateRegex(lastName, pattern)
}

func IsValidLastName(fl validator.FieldLevel) bool {
	lastName := fl.Field().String()
	return validateLastName(lastName)
}
