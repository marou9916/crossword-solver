package solver

import (
	"crossword-solver/internal/model"
	"reflect"
	"testing"
)

func TestFillGridWithContent(t *testing.T) {
	tests := []struct {
		name         string
		grid         [][]rune
		positions    []model.Position
		content      rune
		expectedGrid [][]rune
	}{
		{
			name: "Test 1",
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
			name: "Test 2",
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
		t.Run(test.name, func(t *testing.T) {
			FillGridWithContent(test.grid, test.positions, test.content)
			if !reflect.DeepEqual(test.grid, test.expectedGrid) {
				t.Errorf("FillGridWithContent(%v, %v, %q) = %v; want %v", test.grid, test.positions, test.content, test.grid, test.expectedGrid)
			}
		})
	}
}
