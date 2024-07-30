package solver

import (
	"reflect"
	"testing"
	"crossword-solver/internal/model"
)

func TestGetDotPositions(t *testing.T) {
	grid := [][]rune{
		{'1', '.', '1'},
		{'0', '.', '0'},
		{'.', '0', '1'},
	}

	expected := []model.Position{
		{Row: 0, Col: 1},
		{Row: 1, Col: 1},
		{Row: 2, Col: 0},
	}

	result := GetDotPositions(grid)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetDotPositions(%v) = %v; want %v", grid, result, expected)
	}
}
