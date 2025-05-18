package helper

import (
	"errors"
	"path/filepath"
	"testing"
)

func TestCheckExtension(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      error
	}{
		{".jpg", ".jpg", nil},
		{".png", ".png", nil},
		{".jpeg", ".jpeg", nil},
		{".gif", "", errors.New("type of image should be jpg, png, jpeg")},
	}

	for _, test := range tests {
		result, err := CheckExtension(test.input)
		if result != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
		if (err != nil && test.err == nil) || (err == nil && test.err != nil) {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}

func TestGenerateDst(t *testing.T) {
	path := "uploads/images"
	file := "pic.jpg"
	expected := filepath.Join(path, file)
	result := GenerateDst(path, file)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
