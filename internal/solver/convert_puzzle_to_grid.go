package solver

import (
	"strings"
)

func ConvertPuzzleToGrid(puzzle string) [][]rune {
	var grid [][]rune

	// Diviser le puzzle en lignes
	lines := strings.Split(puzzle, "\n")

	for _, line := range lines {
		if line != "" { // Ignorer les lignes vides
			grid = append(grid, []rune(line))
		}
	}

	return grid
}
