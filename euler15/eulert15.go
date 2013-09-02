package main

import "fmt"

type Point struct {
	X, Y int
}

const (
	gridSize = 21
)

func main() {
	grid := make([][]int, gridSize)
	for i := 0; i < gridSize; i++ {
		grid[i] = make([]int, gridSize)
		grid[i][gridSize-1] = 1
	}
	for idx := range grid[gridSize-1] {
		grid[gridSize-1][idx] = 1
	}
	for col := gridSize - 2; col >= 0; col-- {
		for row := gridSize - 2; row >= 0; row-- {
			for idx := row; idx <= gridSize-1; idx++ {
				grid[col][row] += grid[col+1][idx]
			}
		}
	}
	fmt.Println(grid[0][0])
}
