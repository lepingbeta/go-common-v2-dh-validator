/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-08 20:35:00
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-07-05 04:35:32
 * @FilePath     : /v2/go-common-v2-dh-validator/msg_key_manual.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package dhvalidator

import (
	"github.com/go-playground/validator/v10"
)

func CustomErrors(fe validator.FieldError) string {

	// parts := strings.Fields(fe.Param())
	// 通过判断错误的标签来返回不同的错误消息
	switch fe.Tag() {
	// case "findInDb":
	// 	return parts[3]
	default:
		return GetMsgKey(fe.Tag())
	}
}
