package dhvalidator

import "testing"

// 测试validatePassword函数
// TestValidatePassword 测试validatePassword函数
func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     bool
	}{
		{"short", "abc", false},
		{"long", "aVeryLongPasswordThatExceedsTheMaximumAllowedLength", false},
		{"validLength", "correctLength", true},
		{"withSpecialChars", "valid!Password@123", false},
		{"withSpaces", "valid Password", false},
		{"empty", "", false},
		{"exactly6", "abc123", true},
		{"exactly30", "123456789012345678901234567890", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validatePassword(tt.password); got != tt.want {
				t.Errorf("validatePassword(%q) = %v, want %v", tt.password, got, tt.want)
			}
		})
	}
}
