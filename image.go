package dhvalidator

import (
	"github.com/go-playground/validator/v10"
)

func validateImage(image string) bool {
	// 定义正则表达式
	pattern := "^[a-zA-Z0-9_-]+(\\.(jpg|jpeg|png|gif|webp))$"

	return validateRegex(image, pattern)
}

func IsValidImage(fl validator.FieldLevel) bool {
	image := fl.Field().String()
	return validateImage(image)
}
