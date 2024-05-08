package dhvalidator

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func CustomErrors(fe validator.FieldError) string {

	parts := strings.Fields(fe.Param())
	// 通过判断错误的标签来返回不同的错误消息
	switch fe.Tag() {
	case "findInDb":
		return parts[3]
	default:
		return GetMsgKey(fe.Tag())
	}
}
