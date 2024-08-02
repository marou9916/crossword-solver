package remover

import (
	"crossword-solver/internal/model"
	"reflect"
	"testing"
)

func TestRemoveWordHorizontally(t *testing.T) {
	tests := []struct {
		grid         [][]rune
		word         string
		pos          model.Position
		expectedGrid [][]rune
	}{
		{
			grid: [][]rune{
				{'c', 'a', 't'},
				{' ', ' ', ' '},
				{' ', ' ', ' '},
			},
			word: "cat",
			pos:  model.Position{Row: 0, Col: 0},
			expectedGrid: [][]rune{
				{'0', '0', '0'},
				{' ', ' ', ' '},
				{' ', ' ', ' '},
			},
		},
		{
			grid: [][]rune{
				{' ', ' ', ' '},
				{' ', ' ', ' '},
				{'d', 'o', 'g'},
			},
			word: "dog",
			pos:  model.Position{Row: 2, Col: 0},
			expectedGrid: [][]rune{
				{' ', ' ', ' '},
				{' ', ' ', ' '},
				{'0', '0', '0'},
			},
		},
	}

	for _, test := range tests {
		RemoveWordHorizontally(test.grid, test.word, test.pos)
		if !reflect.DeepEqual(test.grid, test.expectedGrid) {
			t.Errorf("RemoveWordHorizontally(%v, %s, %v) = %v; want %v", test.grid, test.word, test.pos, test.grid, test.expectedGrid)
		}
	}
}
