/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-08 20:35:00
 * @LastEditors: Symphony zhangleping@cezhiqiu.com
 * @LastEditTime: 2024-09-05 22:15:07
 * @FilePath     : /v2/go-common-v2-dh-validator/msg_key.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package dhvalidator

func GetMsgKey(validName string) string {
	ErrorList := map[string]string{
		"firstName":     "validator_invalid_firstname",
		"chineseMobile": "validator_invalid_chinese_mobile",
		"inviteCodeNum": "validator_invalid_invite_code",
		"lastName":      "validator_invalid_lastname",
		"nickname":      "validator_invalid_nickname",
		"gender":        "validator_invalid_gender",
		"mongoId":       "validator_invalid_mongoid",
		"image":         "validator_invalid_image",
		"account":       "validator_invalid_account",
		"password":      "validator_invalid_password",
		"findInDb":      "validator_invalid_findInDb",
		"pascalCase":    "validator_invalid_pascalCase",
		"snakeCase":     "validator_invalid_snakeCase",
	}
	return ErrorList[validName]
}
