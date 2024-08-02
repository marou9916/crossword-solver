// cmd/crosswordsolver/main.go
package main

import (
	"crossword-solver/internal/solver"
	"crossword-solver/internal/validator"
	"fmt"
)

func main() {
	puzzle := `2001
0..0
1000
0..0`
	words := []string{"casa", "alan", "ciao", "anta"}

	// Validation des entrées
	if !validator.IsValidString(puzzle) || !validator.IsValidSetOfWords(words) {
		fmt.Println("Error")
		return
	}

	// Conversion de la grille
	grid := solver.ConvertPuzzleToGrid(puzzle)

	// Récupération des positions
	dotPositions := solver.GetDotPositions(grid)
	wordPositions := solver.GetPositionsForWords(grid)

	// Initialisation de la grille de solutions
	solutionGrid := make([][]rune, len(grid))
	for i := range grid {
		solutionGrid[i] = make([]rune, len(grid[i]))
		copy(solutionGrid[i], grid[i])
	}
	solver.FillGridWithContent(solutionGrid, dotPositions, '.')

	// Backtracking pour trouver une solution
	if solver.TryToFindSolutionWithBacktracking(solutionGrid, words, wordPositions, 0) {
		fmt.Println("Solution found:")
		for _, row := range solutionGrid {
			fmt.Println(string(row))
		}
	} else {
		fmt.Println("No solution found")
	}
}
