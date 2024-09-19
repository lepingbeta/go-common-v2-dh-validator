/*
 * @Author: Symphony zhangleping@cezhiqiu.com
 * @Date: 2024-09-13 19:49:21
 * @LastEditors: Symphony zhangleping@cezhiqiu.com
 * @LastEditTime: 2024-09-13 19:49:26
 * @FilePath: /go-common/v2/go-common-v2-dh-validator/pascalCase.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dhvalidator

import (
	"github.com/go-playground/validator/v10"
)

func validatePascalCase(pascalCase string) bool {
	// 定义正则表达式
	pattern := `^[A-Z][a-z0-9]+(?:[A-Z][a-z0-9]+)*$`

	return validateRegex(pascalCase, pattern)
}

func IsValidPascalCase(fl validator.FieldLevel) bool {
	pascalCase := fl.Field().String()
	return validatePascalCase(pascalCase)
}
