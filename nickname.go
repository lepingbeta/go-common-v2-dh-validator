package dhvalidator

import (
	"github.com/go-playground/validator/v10"
)


func validateNickname(nickname string) bool {
	// 定义正则表达式
	pattern := "^[\\p{L}][\\p{L}\\p{N}_-]{2,19}$"

	return validateRegex(nickname, pattern)
}

func IsValidNickname(fl validator.FieldLevel) bool {
	nickname := fl.Field().String()
	return validateNickname(nickname)
}
