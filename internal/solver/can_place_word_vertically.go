package solver

import "crossword-solver/internal/model"

func CanPlaceWordVertically(grid [][]rune, word string, pos model.Position) bool {
	wordLen := len(word)
	gridHeight := len(grid)

	// Condition 1: Débordement de la grille
	if pos.Row+wordLen > gridHeight {
		return false
	}

	for i := 0; i < wordLen; i++ {
		currentCell := grid[pos.Row+i][pos.Col]

		// Conditions 2 et 3 combinées: Conflit avec des lettres existantes ou cellules bloquantes ou blocage par des cellules bloquantes
		if currentCell != '0' && currentCell != rune(word[i]) || currentCell == '.' {
			return false
		}
	}

	return true
}
