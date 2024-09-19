/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-08 20:35:00
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-09-19 18:10:57
 * @FilePath     : /apibot-api/data/mycode/dahe/go-common/v2/go-common-v2-dh-validator/msg_key_manual.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package dhvalidator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
)

func CustomErrors(fe validator.FieldError) string {
	// dhlog.DebugAny(fe.Param())
	// dhlog.DebugAny(fe.Field())
	// dhlog.DebugAny(fe.StructField())
	// dhlog.DebugAny(fe.Type().Align())
	parts := strings.Fields(fe.Param())
	// 通过判断错误的标签来返回不同的错误消息
	dhlog.DebugAny(fe.Tag())
	switch fe.Tag() {
	case "findInDb":
		return parts[3]
	default:
		return GetMsgKey(fe.Tag())
	}
}

func CustomErrors2(module string, apiName string, fe validator.FieldError) string {
	// dhlog.DebugAny(fe.Param())
	// dhlog.DebugAny(fe.Field())
	// dhlog.DebugAny(fe.Tag())
	formattedString := fmt.Sprintf("param_%v_%v_%v_%v", module, apiName, fe.Field(), fe.Tag())
	return formattedString
	// parts := strings.Fields(fe.Param())
	// // 通过判断错误的标签来返回不同的错误消息

	// switch fe.Tag() {
	// case "findInDb":
	// 	return parts[3]
	// default:
	// 	return GetMsgKey(fe.Tag())
	// }
}
