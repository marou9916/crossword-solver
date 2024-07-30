package solver

import (
	"crossword-solver/internal/model"
)

func GetDotPositions(grid [][]rune) []model.Position {
    return FindPositions(grid, []rune{'.'})
}
