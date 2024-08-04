package remover

import "crossword-solver/internal/model"

func RemoveWordVertically(grid [][]rune, word string, pos model.Position) {
	for i := range word {
		grid[pos.Row+i][pos.Col] = '0'
	}
}
