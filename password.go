package dhvalidator

import (
	"github.com/go-playground/validator/v10"
)

func validatePassword(password string) bool {
	// 定义正则表达式
	pattern := "^[a-zA-Z0-9]{6,30}$"

	return validateRegex(password, pattern)
}

func IsValidPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	return validatePassword(password)
}
