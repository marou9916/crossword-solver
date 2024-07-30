package solver

import (
	"crossword-solver/internal/model"
)

func FillGridWithContent(grid [][]rune, positions []model.Position, content rune) {
	for _, position := range positions {
		if position.Row >= 0 && position.Row < len(grid) && position.Col >= 0 && position.Col < len(grid[position.Row]) {
			grid[position.Row][position.Col] = content
		}
	}
}
