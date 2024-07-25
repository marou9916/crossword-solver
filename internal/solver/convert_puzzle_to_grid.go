package solver

import (
	"fmt"
	"strings"
)

func ConvertPuzzleToGrid(puzzle string) [][]rune {
	var grid [][]rune

	// Diviser le puzzle en lignes
	lines := strings.Split(puzzle, "\n")
	fmt.Println(lines)

	for _, line := range lines {
		if line != "" { // Ignorer les lignes vides
			grid = append(grid, []rune(line))
		}
	}

	return grid
}
