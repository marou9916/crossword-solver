package placement

import (
	"crossword-solver/internal/model"
	"crossword-solver/internal/printer"
)

func PlaceWordVertically(grid [][]rune, word string, pos model.Position) {
	for i, letter := range word {
		grid[pos.Row+i][pos.Col] = letter
	}
	printer.PrintGrid(grid)
}
