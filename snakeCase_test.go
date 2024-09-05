/*
 * @Author: Symphony zhangleping@cezhiqiu.com
 * @Date: 2024-09-05 21:27:32
 * @LastEditors: Symphony zhangleping@cezhiqiu.com
 * @LastEditTime: 2024-09-05 22:11:55
 * @FilePath: /go-common/v2/go-common-v2-dh-validator/isPascalCase.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dhvalidator

import "testing"

func Test_validateSnakeCase(t *testing.T) {
	// 定义测试用例
	tests := []struct {
		input    string
		expected bool
	}{
		{"valid_snake_case", true},
		{"valid123_snake_case", true},
		{"invalid-snake-case", false},    // 包含连字符
		{"Invalid_Snake_Case", false},    // 包含大写字母
		{"snake_case_", false},           // 下划线结尾
		{"_snake_case", false},           // 下划线开头
		{"snake__case", false},           // 连续下划线
		{"", false},                      // 空字符串
		{"valid_snake_case_1_2_3", true}, // 混合字母与数字
	}

	// 遍历测试用例并执行
	for _, test := range tests {
		result := validateSnakeCase(test.input)
		if result != test.expected {
			t.Errorf("validateSnakeCase(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}
