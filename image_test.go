package dhvalidator

import "testing"

// TestValidateImage 测试 validateImage 函数
func TestValidateImage(t *testing.T) {
	tests := []struct {
		name     string // 测试用例的名称
		filename string // 测试用例的输入值
		want     bool   // 期望的验证结果
	}{
		{"ValidJPG", "image.jpg", true},
		{"ValidJPEG", "photo.jpeg", true},
		{"ValidPNG", "picture.png", true},
		{"ValidGIF", "animation.gif", true},
		{"ValidWebP", "graphic.webp", true},
		{"InvalidExtension", "document.pdf", false},
		{"NoExtension", "image", false},
		{"InvalidCharacter", "image*.jpg", false},
		{"LeadingSpace", " image.jpg", false},
		{"TrailingSpace", "image.jpg ", false},
		{"UpperCaseExtension", "photo.JPG", false}, // 如果你的系统区分文件名大小写，这可能需要调整
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validateImage(tt.filename)
			if got != tt.want {
				t.Errorf("%s: validateImage(%v) = %v, want %v", tt.name, tt.filename, got, tt.want)
			}
		})
	}
}
