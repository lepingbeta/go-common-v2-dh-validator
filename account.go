package dhvalidator

import (
	"github.com/go-playground/validator/v10"
)

func validateAccount(account string) bool {
	// 定义正则表达式
	pattern := "^[a-zA-Z][a-zA-Z0-9_]{4,30}$"

	return validateRegex(account, pattern)
}

func IsValidAccount(fl validator.FieldLevel) bool {
	account := fl.Field().String()
	return validateAccount(account)
}
