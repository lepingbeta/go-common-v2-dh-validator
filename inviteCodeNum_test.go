/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-05 10:47:49
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-05-05 10:56:43
 * @FilePath     : /dahe/go-common/dh-validator/inviteCodeNum_test.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package dhvalidator

import "testing"

func Test_validateInviteCodeNum(t *testing.T) {
	// 测试用例，包括预期结果
	tests := []struct {
		inviteCodeNum string
		expected      bool
	}{
		{"12345", true},   // 正确的5位数字邀请码
		{"1234", false},   // 少于5位数字
		{"123456", false}, // 多于5位数字
		{"abcde", false},  // 包含非数字字符
		{"", false},       // 空字符串
	}

	for _, tt := range tests {
		if got := validateInviteCodeNum(tt.inviteCodeNum); got != tt.expected {
			t.Errorf("validateInviteCodeNum(%q) = %v, want %v", tt.inviteCodeNum, got, tt.expected)
		}
	}
}
