package solver

import (
	"reflect"
	"testing"
)

func TestConvertPuzzleToGrid(t *testing.T) {
	puzzle := "2001\n0..0\n1000\n0..0\n"
	expected := [][]rune{
		{'2', '0', '0', '1'},
		{'0', '.', '.', '0'},
		{'1', '0', '0', '0'},
		{'0', '.', '.', '0'},
	}

	result := ConvertPuzzleToGrid(puzzle)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("convertPuzzleToGrid(%q) = %v; want %v", puzzle, result, expected)
	}
}
