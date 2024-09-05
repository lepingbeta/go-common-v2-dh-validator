/*
 * @Author: Symphony zhangleping@cezhiqiu.com
 * @Date: 2024-09-05 21:27:32
 * @LastEditors: Symphony zhangleping@cezhiqiu.com
 * @LastEditTime: 2024-09-05 21:30:26
 * @FilePath: /go-common/v2/go-common-v2-dh-validator/isPascalCase.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dhvalidator

import (
	"github.com/go-playground/validator/v10"
)

func validatePascalCase(account string) bool {
	// 定义正则表达式
	pattern := "^[A-Z][a-z0-9]+(?:[A-Z][a-z0-9]+)*$"

	return validateRegex(account, pattern)
}

func IsValidPascalCase(fl validator.FieldLevel) bool {
	account := fl.Field().String()
	return validatePascalCase(account)
}
