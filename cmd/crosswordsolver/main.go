// cmd/crosswordsolver/main.go
package main

import (
	"crossword-solver/internal/printer"
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
		fmt.Println("Error in validation")
		return
	}

	// Conversion de la grille
	grid := solver.ConvertPuzzleToGrid(puzzle)
	fmt.Println("Initial Grid:")
	printer.PrintGrid(grid)

	// Récupération des positions
	dotPositions := solver.GetDotPositions(grid)
	wordPositions := solver.GetPositionsForWords(grid)
	fmt.Println("Dot Positions:", dotPositions)
	fmt.Println("Word Positions:", wordPositions)

	// Initialisation de la grille de solutions
	solutionGrid := make([][]rune, len(grid))
	for i := range grid {
		solutionGrid[i] = make([]rune, len(grid[i]))
		for j := range solutionGrid[i] {
			solutionGrid[i][j] = ' ' // Initialise avec des points
		}
	}

	solver.FillGridWithContent(solutionGrid, dotPositions, '.')
	fmt.Println("Initial Solution Grid:")
	printer.PrintGrid(solutionGrid)

	// Backtracking pour trouver une solution
	if solver.TryToFindSolutionWithBacktracking(solutionGrid, words, wordPositions, 0) {
		fmt.Println("Solution found:")
		printer.PrintGrid(solutionGrid)
	} else {
		fmt.Println("No solution found")
	}
}
