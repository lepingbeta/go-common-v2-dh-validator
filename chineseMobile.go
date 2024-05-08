package dhvalidator

import (
	"github.com/go-playground/validator/v10"
)


func validateChineseMobile(chineseMobile string) bool {
	// 定义正则表达式
	pattern := "^1[3-9]\\d{9}$"

	return validateRegex(chineseMobile, pattern)
}

func IsValidChineseMobile(fl validator.FieldLevel) bool {
	chineseMobile := fl.Field().String()
	return validateChineseMobile(chineseMobile)
}
