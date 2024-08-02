package remover

import "crossword-solver/internal/model"

func RemoveWordHorizontally(grid [][]rune, word string, pos model.Position) {
	for i := range word {
		grid[pos.Row][pos.Col+i] = '0'
	}
}
