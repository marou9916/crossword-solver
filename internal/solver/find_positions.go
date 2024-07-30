package solver

import (
	"crossword-solver/internal/model"
)

func FindPositions(grid [][]rune, searchedContents []rune) []model.Position {
	var positions []model.Position

	for rowIndex, line := range grid {
		for colIndex, contentOfTheLine := range line {
			for _, content := range searchedContents {
				if contentOfTheLine == content {
					positions = append(positions, model.Position{Row: rowIndex, Col: colIndex})
				}
			}
		}
	}

	return positions
}
