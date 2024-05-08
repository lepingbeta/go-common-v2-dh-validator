package dhvalidator

import "testing"

// TestValidateChineseMobile 测试 ValidateChineseMobile 函数
func TestValidateChineseMobile(t *testing.T) {
	// 定义测试案例
	tests := []struct {
		name   string // 测试案例的名称
		mobile string // 输入的手机号
		want   bool   // 预期的验证结果
	}{
		{"ValidMobile", "13912345678", true},
		{"InvalidMobileShort", "139123456", false},
		{"InvalidMobileLong", "139123456789", false},
		{"InvalidMobileWithLetters", "1391234567a", false},
		{"InvalidMobileWithSpecialChars", "13912345_78", false},
		{"InvalidMobileStartWith2", "23912345678", false},
		{"ValidMobileWith188", "18812345678", true},
		{"ValidMobileWith199", "19912345678", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validateChineseMobile(tt.mobile)
			if got != tt.want {
				t.Errorf("ValidateChineseMobile(%v) = %v, want %v", tt.mobile, got, tt.want)
			}
		})
	}
}
