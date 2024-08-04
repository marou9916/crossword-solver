package placement_test

import (
	"crossword-solver/internal/model"
	"crossword-solver/internal/placement"
	"reflect"
	"testing"
)

func TestPlaceWordVertically(t *testing.T) {
	tests := []struct {
		grid     [][]rune
		word     string
		position model.Position
		expected [][]rune
	}{
		{
			grid: [][]rune{
				{' ', ' ', ' '},
				{' ', ' ', ' '},
				{' ', ' ', ' '},
			},
			word:     "ABC",
			position: model.Position{Row: 0, Col: 0},
			expected: [][]rune{
				{'A', ' ', ' '},
				{'B', ' ', ' '},
				{'C', ' ', ' '},
			},
		},
	}

	for _, test := range tests {
		placement.PlaceWordVertically(test.grid, test.word, test.position)
		if !reflect.DeepEqual(test.grid, test.expected) {
			t.Errorf("PlaceWordVertically(%v, %s, %v) = %v; want %v", test.grid, test.word, test.position, test.grid, test.expected)
		}
	}
}
