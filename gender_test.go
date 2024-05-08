package dhvalidator

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestIsValidGender(t *testing.T) {
	tests := []struct {
		gender string
		valid  bool
	}{
		{"male", true},
		{"female", true},
		{"other", true},
		{"invalid_gender", false},
		{"Male", false},
		{"Female", false},
	}

	for _, tt := range tests {
		t.Run(tt.gender, func(t *testing.T) {
			assert.Equal(t, tt.valid, validateGender(tt.gender))
		})
	}
}
