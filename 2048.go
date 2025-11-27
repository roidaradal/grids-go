package main

import (
	"fmt"

	"github.com/roidaradal/fn/lang"
)

// 2048: 4x4 grid, adds 2 randomly at empty space, with small chance of 4

// Create new 2048 grid
func new2048() Grid[int] {
	// Make blank int grid
	numRows, numCols := 4, 4
	grid := newGrid[int](numRows, numCols)
	// Randomly place 2x 2 tiles
	for _, idx := range randomNumbers(numRows*numCols, 2) {
		row, col := indexToCoords(idx, numCols).Tuple()
		grid[row][col] = 2
	}
	return grid
}

// Display 2048 grid
func display2048(grid Grid[int]) {
	blankCell := fmt.Sprintf("%5s", ".")
	displayGrid(grid, func(cell int) string {
		numCell := fmt.Sprintf("%5d", cell)
		return lang.Ternary(cell == 0, blankCell, numCell)
	})
}
