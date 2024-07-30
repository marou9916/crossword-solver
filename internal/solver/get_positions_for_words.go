package solver

import (
	"crossword-solver/internal/model"
)

func GetPositionsForWords(grid [][]rune) []model.Position {
	searchedContents := []rune{'1', '2'}
	return FindPositions(grid, searchedContents)
}
