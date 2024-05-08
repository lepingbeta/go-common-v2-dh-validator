/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-05 10:47:49
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-05-05 10:53:37
 * @FilePath     : /dahe/go-common/dh-validator/inviteCodeNum.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package dhvalidator

import (
	"github.com/go-playground/validator/v10"
)

func validateInviteCodeNum(inviteCodeNum string) bool {
	// 定义正则表达式
	pattern := `^\d{5}$`

	return validateRegex(inviteCodeNum, pattern)
}

func IsValidInviteCodeNum(fl validator.FieldLevel) bool {
	inviteCodeNum := fl.Field().String()
	return validateInviteCodeNum(inviteCodeNum)
}
