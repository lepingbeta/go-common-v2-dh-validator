package dhvalidator

import "testing"

// TestValidateNickname 测试 validateNickname 函数
func TestValidateNickname(t *testing.T) {
	// 定义测试案例
	tests := []struct {
		name     string // 测试案例的名称
		nickname string // 输入的昵称
		want     bool   // 预期的验证结果
	}{
		{"ValidSingleWord", "Nickname", true},
		{"ValidChinese", "昵称123", true},
		{"ValidWithNumbers", "Nick123", true},
		{"ValidWithUnderscore", "Nick_Name", true},
		{"ValidWithHyphen", "Nick-Name", true},
		{"ValidJapanese", "ニックネーム", true},
		{"InvalidStartWithNumber", "123Nickname", false},
		{"InvalidCharacters", "Nick*Name", false},
		{"TooShort", "Ni", false},
		{"TooLong", "N123456789012345678901234567890", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validateNickname(tt.nickname)
			if got != tt.want {
				t.Errorf("validateNickname(%v) = %v, want %v", tt.nickname, got, tt.want)
			}
		})
	}
}
