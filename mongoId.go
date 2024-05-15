package dhvalidator

import (
	"github.com/go-playground/validator/v10"
)


func validateMongoId(mongoId string) bool {
	// 定义正则表达式
	pattern := `^[0-9a-fA-F]{24}$`

	return validateRegex(mongoId, pattern)
}

func IsValidMongoId(fl validator.FieldLevel) bool {
	mongoId := fl.Field().String()
	return validateMongoId(mongoId)
}
