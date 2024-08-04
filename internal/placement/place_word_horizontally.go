package placement

import (
	"crossword-solver/internal/model"
	"crossword-solver/internal/printer"
)

func PlaceWordHorizontally(grid [][]rune, word string, pos model.Position) {
	for i, letter := range word {
		grid[pos.Row][pos.Col+i] = letter
	}
	printer.PrintGrid(grid)
}
