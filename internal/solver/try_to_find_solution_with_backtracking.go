package solver

import (
	"crossword-solver/internal/model"
	"crossword-solver/internal/placement"
	"crossword-solver/internal/remover"
)

func TryToFindSolutionWithBacktracking(grid [][]rune, words []string, positions []model.Position, wordIndex int) bool {

	if wordIndex >= len(words) {
		return true
	}

	word := words[wordIndex]

	for _, position := range positions {
		if placement.CanPlaceWordHorizontally(grid, word, position) {
			placement.PlaceWordHorizontally(grid, word, position)
			if TryToFindSolutionWithBacktracking(grid, words, positions, wordIndex+1) {
				return true
			}
			remover.RemoveWordHorizontally(grid, word, position)
		}

		if placement.CanPlaceWordVertically(grid, word, position) {
			placement.PlaceWordVertically(grid, word, position)
			if TryToFindSolutionWithBacktracking(grid, words, positions, wordIndex+1) {
				return true
			}
			remover.RemoveWordVertically(grid, word, position)
		}
	}

	return false
}
