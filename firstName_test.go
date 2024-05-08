package dhvalidator

import "testing"

// TestValidateFirstName 测试 validateFirstName 函数。
func TestValidateFirstName(t *testing.T) {
	tests := []struct {
		name      string // 测试用例的名称
		firstName string // 测试用例的输入值
		want      bool   // 期望的验证结果
	}{
		{"ValidName", "John", true},
		{"ValidNameWithAccent", "José", true},
		{"ValidChineseName", "王伟", true},
		{"ValidNameWithNumber", "Alice1", true},
		{"ValidNameWithUnderscore", "Bob_TheBuilder", true},
		{"ValidNameWithHyphen", "Jean-Luc", true},
		{"EmptyName", "", false},
		{"NameTooLong", "ThisNameIsWayTooLongForTheValidator", false},
		{"InvalidNameWithSpecialChars", "Anna*Marie", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validateFirstName(tt.firstName)
			if got != tt.want {
				t.Errorf("%s: validateFirstName(%v) = %v, want %v", tt.name, tt.firstName, got, tt.want)
			}
		})
	}
}
