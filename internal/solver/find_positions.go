package solver

func FindPositions(grid [][]rune, searchedContent rune) []Position {
	var positionsWhereContentIsFound []Position

	for rowIndex, line := range grid {
		for colIndex, contentOfTheLine := range line {
			if contentOfTheLine == searchedContent {
				position := Position {row : rowIndex, col : colIndex}
				positionsWhereContentIsFound = append(positionsWhereContentIsFound, position)
			}
		}
	}

	return positionsWhereContentIsFound
}
