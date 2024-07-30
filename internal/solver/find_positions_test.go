package solver

import (
	"crossword-solver/internal/model"
	"reflect"
	"testing"
)

func TestFindPositions(t *testing.T) {
	tests := []struct {
		name             string
		grid             [][]rune
		searchedContents []rune
		expected         []model.Position
	}{
		{
			name: "Find positions of '1' and '2'",
			grid: [][]rune{
				{'1', '0', '1'},
				{'0', '2', '0'},
				{'1', '0', '1'},
			},
			searchedContents: []rune{'1', '2'},
			expected: []model.Position{
				{Row: 0, Col: 0}, {Row: 0, Col: 2},
				{Row: 1, Col: 1},
				{Row: 2, Col: 0}, {Row: 2, Col: 2},
			},
		},
		{
			name: "Find positions of '.'",
			grid: [][]rune{
				{'1', '.', '1'},
				{'0', '.', '0'},
				{'.', '0', '1'},
			},
			searchedContents: []rune{'.'},
			expected: []model.Position{
				{Row: 0, Col: 1},
				{Row: 1, Col: 1},
				{Row: 2, Col: 0},
			},
		},
		{
			name: "Mixed content",
			grid: [][]rune{
				{'1', '0', '1'},
				{'2', '.', '2'},
				{'0', '0', '0'},
			},
			searchedContents: []rune{'1', '.', '2'},
			expected: []model.Position{
				{Row: 0, Col: 0}, {Row: 0, Col: 2},
				{Row: 1, Col: 0}, {Row: 1, Col: 1}, {Row: 1, Col: 2},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FindPositions(test.grid, test.searchedContents)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("FindPositions(%v, %v) = %v; want %v", test.grid, test.searchedContents, result, test.expected)
			}
		})
	}
}
