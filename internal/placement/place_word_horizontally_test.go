package placement

import (
	"crossword-solver/internal/model"
	"reflect"
	"testing"
)

func TestPlaceWordHorizontally(t *testing.T) {
	tests := []struct {
		name         string
		grid         [][]rune
		word         string
		pos          model.Position
		expectedGrid [][]rune
	}{
		{
			name: "Test 1",
			grid: [][]rune{
				{' ', ' ', 't'},
				{' ', ' ', 'a'},
				{' ', ' ', 'g'},
			},
			word: "cat",
			pos:  model.Position{Row: 0, Col: 0},
			expectedGrid: [][]rune{
				{'c', 'a', 't'},
				{' ', ' ', 'a'},
				{' ', ' ', 'g'},
			},
		},
		{
			name: "Test 2",
			grid: [][]rune{
				{' ', ' ', ' '},
				{' ', ' ', ' '},
				{' ', ' ', ' '},
			},
			word: "dog",
			pos:  model.Position{Row: 2, Col: 0},
			expectedGrid: [][]rune{
				{' ', ' ', ' '},
				{' ', ' ', ' '},
				{'d', 'o', 'g'},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			PlaceWordHorizontally(test.grid, test.word, test.pos)
			if !reflect.DeepEqual(test.grid, test.expectedGrid) {
				t.Errorf("PlaceWordHorizontally(%v, %s, %v) = %v; want %v", test.grid, test.word, test.pos, test.grid, test.expectedGrid)
			}
		})
	}
}
