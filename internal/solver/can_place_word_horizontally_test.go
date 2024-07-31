package solver

import (
	"crossword-solver/internal/model"
	"testing"
)

func TestCanPlaceWordHorizontally(t *testing.T) {
	grid := [][]rune{
		{'c', '0', '0', '0'},
		{'i', '.', '.', '0'},
		{'a', '0', '0', '0'},
		{'o', '.', '.', '0'},
	}

	tests := []struct {
		word     string
		pos      model.Position
		expected bool
	}{
		{"word", model.Position{Row: 0, Col: 0}, false},
		{"wor", model.Position{Row: 0, Col: 2}, false},
		{"wor", model.Position{Row: 2, Col: 1}, true},
		{"w", model.Position{Row: 3, Col: 1}, false},
		// Ajout de tests pour les lettres correspondantes
		{"anta", model.Position{Row: 2, Col: 0}, true},
	}

	for _, test := range tests {
		result := CanPlaceWordHorizontally(grid, test.word, test.pos)
		if result != test.expected {
			t.Errorf("CanPlaceWordHorizontally(%q, %v) = %v; want %v", test.word, test.pos, result, test.expected)
		}
	}
}
