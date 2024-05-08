package dhvalidator

import (
	"testing"
)

// TestValidateAccount 测试 validateAccount 函数。
func TestValidateAccount(t *testing.T) {
	// 定义测试案例
	tests := []struct {
		name    string // 测试案例名称
		account string // 输入的账号
		want    bool   // 期望的结果
	}{
		{"validSimple", "abc123", true},
		{"validWithUnderscore", "abc_123", true},
		{"startWithNumber", "1abc123", false},
		{"startWithUnderscore", "_abc123", false},
		{"tooShort", "a1", false},
		{"tooLong", "a1234567890123456789012345678901", false},
		{"empty", "", false},
		{"exactly6", "abc123", true},
		{"exactly30", "a12345678901234567890123456789", true},
		{"validAllUnderscore", "a______", true},
		{"validComplex", "Aa1_123", true},
		{"validComplex", "admin", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateAccount(tt.account); got != tt.want {
				t.Errorf("validateAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
