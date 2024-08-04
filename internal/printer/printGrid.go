package printer

import "fmt"

func PrintGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
	fmt.Println()
}
