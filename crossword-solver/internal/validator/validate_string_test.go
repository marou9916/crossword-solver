package validator

import (
	"testing"
)

func TestIsValidString(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"2001\n0..0\n1000\n0..0\n", false},    // Chaîne qui se termine par un saut de ligne
		{"2001\n0..0\n1000\n0..0", true},       // Chaîne valide
		{"2001\n0..0\n100\n0..0", false},       // Chaîne avec une ligne de longueur différente
		{"0..0\n0..0", true},                   // Chaîne valide avec des points
		{"2001\n0..0\n1000\n0..0\n1000", true}, // Chaîne valide
		{"123\n456\n789", false},               // Chaîne invalide
	}

	for _, test := range tests {
		result := IsValidString(test.input)
		if result != test.expected {
			t.Errorf("IsValidString(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

