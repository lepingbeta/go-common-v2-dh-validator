/*
 * @Author: Symphony zhangleping@cezhiqiu.com
 * @Date: 2024-09-05 21:27:32
 * @LastEditors: Symphony zhangleping@cezhiqiu.com
 * @LastEditTime: 2024-09-05 21:33:36
 * @FilePath: /go-common/v2/go-common-v2-dh-validator/isPascalCase.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dhvalidator

import "testing"

func Test_validatePascalCase(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"PascalCase", true},     // 正确的 PascalCase
		{"camelCase", false},     // 小写开头的 CamelCase，不符合 PascalCase
		{"Snake_case", false},    // 带有下划线的 SnakeCase，不符合 PascalCase
		{"Pascalcase", true},     // PascalCase 名称，符合
		{"Pascal1Case", true},    // 含数字，但符合 PascalCase
		{"123Pascal", false},     // 数字开头，不符合 PascalCase
		{"PascalCase123", true},  // 尾部带数字，符合 PascalCase
		{"pascalcase", false},    // 全小写，不符合 PascalCase
		{"PascalcaseTest", true}, // 多个单词，符合 PascalCase
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := validatePascalCase(test.input)
			if result != test.expected {
				t.Errorf("validatePascalCase(%q) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}
