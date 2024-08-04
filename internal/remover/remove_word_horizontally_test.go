package remover

import (
	"crossword-solver/internal/model"
	"reflect"
	"testing"
)

func TestRemoveWordHorizontally(t *testing.T) {
	tests := []struct {
		grid     [][]rune
		word     string
		position model.Position
		expected [][]rune
	}{
		{
			grid: [][]rune{
				{'A', 'B', 'C'},
				{' ', ' ', ' '},
				{' ', ' ', ' '},
			},
			word:     "ABC",
			position: model.Position{Row: 0, Col: 0},
			expected: [][]rune{
				{'0', '0', '0'},
				{' ', ' ', ' '},
				{' ', ' ', ' '},
			},
		},
	}

	for _, test := range tests {
		RemoveWordHorizontally(test.grid, test.word, test.position)
		if !reflect.DeepEqual(test.grid, test.expected) {
			t.Errorf("RemoveWordHorizontally(%v, %s, %v) = %v; want %v", test.grid, test.word, test.position, test.grid, test.expected)
		}
	}
}
