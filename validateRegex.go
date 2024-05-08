/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-08 20:35:01
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-05-08 20:35:32
 * @FilePath     : /v2/go-common-v2-dh-validator/validateRegex.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package dhvalidator

import (
	"regexp"

	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
)

func validateRegex(password, pattern string) bool {

	// 使用正则表达式验证密码
	match, err := regexp.MatchString(pattern, password)
	if err != nil {
		dhlog.Info("Error validating :", err)
		return false
	}
	return match
}
