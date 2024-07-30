package solver

import (
	"reflect"
	"testing"
	"crossword-solver/internal/model"
)


func TestGetPositionsForWords(t *testing.T) {
	grid := [][]rune{
		{'1', '0', '1'},
		{'0', '2', '0'},
		{'1', '0', '1'},
	}

	expected := []model.Position{
		{Row: 0, Col: 0}, {Row: 0, Col: 2},
		{Row: 1, Col: 1},
		{Row: 2, Col: 0}, {Row: 2, Col: 2},
	}

	result := GetPositionsForWords(grid)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetPositionsForWords(%v) = %v; want %v", grid, result, expected)
	}
}
