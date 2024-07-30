package solver

import (
	"reflect"
	"testing"
)

func TestFindPositions(t *testing.T) {
	tests := []struct {
		grid            [][]rune
		searchedContent rune
		expected        []Position
	}{
		{
			grid: [][]rune{
				{'.', 'A', 'B'},
				{'B', '.', 'A'},
				{'A', 'B', '.'},
			},
			searchedContent: 'A',
			expected: []Position{
				{0, 1},
				{1, 2},
				{2, 0},
			},
		},

		{
			grid: [][]rune{
				{'.', '.', '.'},
				{'.', 'A', '.'},
				{'.', '.', '.'},
			},
			searchedContent: 'A',
			expected: []Position{
				{1, 1},
			},
		},
		{
			grid: [][]rune{
				{'A', 'A', '.'},
				{'.', 'A', 'A'},
				{'A', '.', 'A'},
			},
			searchedContent: 'A',
			expected: []Position{
				{0, 0},
				{0, 1},
				{1, 1},
				{1, 2},
				{2, 0},
				{2, 2},
			},
		},
		{
			grid: [][]rune{
				{'#', '#', '#'},
				{'#', '@', '#'},
				{'#', '#', '#'},
			},
			searchedContent: '@',
			expected: []Position{
				{1, 1},
			},
		},
	}

	for _, test := range tests {
		result := FindPositions(test.grid, test.searchedContent)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("GetPositions(%v, %q) = %v; want %v", test.grid, test.searchedContent, result, test.expected)
		}
	}
}
