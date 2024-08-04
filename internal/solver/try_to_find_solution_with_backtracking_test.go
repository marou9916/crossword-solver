package solver

import (
	"crossword-solver/internal/model"
	"reflect"
	"testing"
)

func TestTryToFindSolutionWithBacktracking(t *testing.T) {
	tests := []struct {
		grid           [][]rune
		words          []string
		positions      []model.Position
		expectedGrid   [][]rune
		expectedResult bool
	}{
		{
			grid: [][]rune{
				{'0', '0', '0', '0'},
				{'0', '.', '.', '0'},
				{'0', '0', '0', '0'},
				{'0', '.', '.', '0'},
			},
			words: []string{"AB", "CD"},
			positions: []model.Position{
				{Row: 1, Col: 1},
				{Row: 3, Col: 1},
			},
			expectedGrid: [][]rune{
				{'0', '0', '0', '0'},
				{'0', 'A', 'B', '0'},
				{'0', '0', '0', '0'},
				{'0', 'C', 'D', '0'},
			},
			expectedResult: true,
		},
		{
			grid: [][]rune{
				{'0', '0', '0', '0'},
				{'0', '.', '.', '0'},
				{'0', '0', '0', '0'},
				{'0', '.', '.', '0'},
			},
			words: []string{"ABCDE"},
			positions: []model.Position{
				{Row: 1, Col: 1},
			},
			expectedGrid: [][]rune{
				{'0', '0', '0', '0'},
				{'0', '.', '.', '0'},
				{'0', '0', '0', '0'},
				{'0', '.', '.', '0'},
			},
			expectedResult: false,
		},
	}

	for _, test := range tests {
		result := TryToFindSolutionWithBacktracking(test.grid, test.words, test.positions, 0)
		if result != test.expectedResult || !reflect.DeepEqual(test.grid, test.expectedGrid) {
			t.Errorf("TryToFindSolutionWithBacktracking(%v, %v, %v) = %v, %v; want %v, %v", test.grid, test.words, test.positions, result, test.grid, test.expectedResult, test.expectedGrid)
		}
	}
}
