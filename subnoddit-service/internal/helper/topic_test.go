package helper

import (
	"reflect"
	"testing"
)

func TestSplitTopicIDs(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Basic comma-separated input",
			input:    "id1,id2,id3",
			expected: []string{"id1", "id2", "id3"},
		},
		{
			name:     "Input with spaces",
			input:    " id1 , id2 ,id3 ",
			expected: []string{"id1", "id2", "id3"},
		},
		{
			name:     "Single ID",
			input:    "onlyone",
			expected: []string{"onlyone"},
		},
		{
			name:     "Empty string",
			input:    "",
			expected: []string{""},
		},
		{
			name:     "Trailing comma",
			input:    "id1,id2,",
			expected: []string{"id1", "id2", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SplitTopicIDs(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SplitTopicIDs(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
