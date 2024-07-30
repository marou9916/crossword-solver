package solver

import (
	"reflect"
	"testing"
	"crossword-solver/internal/model"
)
func TestFillGridWithContent(t *testing.T) {
	tests := []struct {
		grid         [][]rune
		positions    []model.Position
		content      rune
		expectedGrid [][]rune
	}{
		{
			grid: [][]rune{
				{' ', ' ', ' '},
				{' ', ' ', ' '},
				{' ', ' ', ' '},
			},
			positions: []model.Position{
				{Row: 0, Col: 0}, {Row: 1, Col: 1}, {Row: 2, Col: 2},
			},
			content: 'X',
			expectedGrid: [][]rune{
				{'X', ' ', ' '},
				{' ', 'X', ' '},
				{' ', ' ', 'X'},
			},
		},
		{
			grid: [][]rune{
				{' ', ' ', ' '},
				{' ', ' ', ' '},
				{' ', ' ', ' '},
			},
			positions: []model.Position{
				{Row: 0, Col: 0}, {Row: 0, Col: 1}, {Row: 0, Col: 2},
			},
			content: 'O',
			expectedGrid: [][]rune{
				{'O', 'O', 'O'},
				{' ', ' ', ' '},
				{' ', ' ', ' '},
			},
		},
	}

	for _, test := range tests {
		FillGridWithContent(test.grid, test.positions, test.content)
		if !reflect.DeepEqual(test.grid, test.expectedGrid) {
			t.Errorf("FillGridWithContent(%v, %v, %q) = %v; want %v", test.grid, test.positions, test.content, test.grid, test.expectedGrid)
		}
	}
}
