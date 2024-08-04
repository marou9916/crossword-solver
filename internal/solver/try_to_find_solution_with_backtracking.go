package solver

import (
	"crossword-solver/internal/model"
	"crossword-solver/internal/placement"
	"crossword-solver/internal/printer"
	"crossword-solver/internal/remover"
	"fmt"
)

func TryToFindSolutionWithBacktracking(grid [][]rune, words []string, positions []model.Position, wordIndex int) bool {
	if wordIndex >= len(words) {
		fmt.Println("All words placed, solution found!")
		return true
	}

	word := words[wordIndex]
	fmt.Printf("Trying to place word %s (Index %d)\n", word, wordIndex)

	for _, position := range positions {
		fmt.Printf("Checking position %v\n", position)
		printer.PrintGrid(grid)

		if placement.CanPlaceWordHorizontally(grid, word, position) {
			fmt.Printf("Placing %s horizontally at %v\n", word, position)
			placement.PlaceWordHorizontally(grid, word, position)

			if TryToFindSolutionWithBacktracking(grid, words, positions, wordIndex+1) {
				return true
			}

			fmt.Printf("Removing %s horizontally from %v\n", word, position)
			remover.RemoveWordHorizontally(grid, word, position)
		}

		if placement.CanPlaceWordVertically(grid, word, position) {
			fmt.Printf("Placing %s vertically at %v\n", word, position)
			placement.PlaceWordVertically(grid, word, position)

			if TryToFindSolutionWithBacktracking(grid, words, positions, wordIndex+1) {
				return true
			}

			fmt.Printf("Removing %s vertically from %v\n", word, position)
			remover.RemoveWordVertically(grid, word, position)
		}
	}

	fmt.Println("Backtracking from word", word)
	return false
}
