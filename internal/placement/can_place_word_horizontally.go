package placement

import "crossword-solver/internal/model"

func CanPlaceWordHorizontally(grid [][]rune, word string, pos model.Position) bool {
	wordLen := len(word)
	gridWidth := len(grid[pos.Row])

	// Condition 1: Débordement de la grille
	if pos.Col+wordLen > gridWidth {
		return false
	}

	for i := 0; i < wordLen; i++ {
		currentCell := grid[pos.Row][pos.Col+i]

		// Conditions 2 et 3 combinées: Conflit avec des lettres existantes ou cellules bloquantes ou blocage par des cellules bloquantes
		if currentCell != '0' && currentCell != rune(word[i]) || currentCell == '.' {
			return false
		}
	}

	return true
}
