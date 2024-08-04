package placement_test

import (
	"crossword-solver/internal/model"
	"crossword-solver/internal/placement"
	"testing"
)

func TestCanPlaceWordVertically(t *testing.T) {
	tests := []struct {
		grid     [][]rune
		word     string
		position model.Position
		expected bool
	}{
		{
			grid: [][]rune{
				{' ', ' ', ' '},
				{' ', ' ', ' '},
				{' ', ' ', ' '},
			},
			word:     "ABC",
			position: model.Position{Row: 0, Col: 0},
			expected: true,
		},
		{
			grid: [][]rune{
				{'A', ' ', ' '},
				{'B', ' ', ' '},
				{'C', ' ', ' '},
			},
			word:     "DEF",
			position: model.Position{Row: 0, Col: 0},
			expected: false,
		},
		{
			grid: [][]rune{
				{' ', ' ', ' '},
				{' ', ' ', ' '},
				{' ', ' ', ' '},
			},
			word:     "ABCDE",
			position: model.Position{Row: 0, Col: 0},
			expected: false,
		},
	}

	for _, test := range tests {
		result := placement.CanPlaceWordVertically(test.grid, test.word, test.position)
		if result != test.expected {
			t.Errorf("CanPlaceWordVertically(%v, %s, %v) = %v; want %v", test.grid, test.word, test.position, result, test.expected)
		}
	}
}
