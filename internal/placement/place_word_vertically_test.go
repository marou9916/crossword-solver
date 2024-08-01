package placement

import (
	"crossword-solver/internal/model"
	"reflect"
	"testing"
)

func TestPlaceWordVertically(t *testing.T) {
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
			word: "tag",
			pos:  model.Position{Row: 0, Col: 2},
			expectedGrid: [][]rune{
				{'c', 'a', 't'},
				{' ', ' ', 'a'},
				{' ', ' ', 'g'},
			},
		},
		{
			grid: [][]rune{
				{' ', ' ', ' '},
				{' ', ' ', ' '},
				{' ', ' ', ' '},
			},
			word: "dog",
			pos:  model.Position{Row: 0, Col: 2},
			expectedGrid: [][]rune{
				{' ', ' ', 'd'},
				{' ', ' ', 'o'},
				{' ', ' ', 'g'},
			},
		},
	}

	for _, test := range tests {
		PlaceWordVertically(test.grid, test.word, test.pos)
		if !reflect.DeepEqual(test.grid, test.expectedGrid) {
			t.Errorf("PlaceWordVertically(%v, %s, %v) = %v; want %v", test.grid, test.word, test.pos, test.grid, test.expectedGrid)
		}
	}
}
