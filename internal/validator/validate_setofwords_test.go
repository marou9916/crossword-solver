package validator

import (
	"testing"
)

func TestIsValidSetOfWords(t *testing.T) {
	tests := []struct {
		input    []string
		expected bool
	}{
		{[]string{"casa", "alan", "ciao", "anta"}, true},
		{[]string{"casa", "alan", "ciao", "alan"}, false},
		{[]string{"", "a", "b", "c"}, true},
		{[]string{"a", "b", "c", "c"}, false},
		{[]string{"unique"}, true},
	}

	for _, test := range tests {
		result := IsValidSetOfWords(test.input)
		if result != test.expected {
			t.Errorf("IsValidSetOfWords(%q) = %v; want %v", test.input, test.expected, result)
		}
	}
}
