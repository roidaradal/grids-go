package main

import (
	"fmt"

	"github.com/roidaradal/fn/list"
)

const greedCursor int = -1

// Create new Greed grid
func newGreed() Grid[int] {
	// Make blank int grid
	numRows, numCols := 22, 79
	grid := newGrid[int](numRows, numCols)
	// Randomize grid digits
	perDigit := 193
	digits := make([]int, 0, numRows*numCols)
	digits = append(digits, greedCursor)
	for i := 1; i <= 9; i++ {
		digits = append(digits, list.Repeated(i, perDigit)...)
	}
	list.Shuffle(digits)
	for idx, digit := range digits {
		row, col := indexToCoords(idx, numCols).Tuple()
		grid[row][col] = digit
	}
	return grid
}

// Display Greed grid
func displayGreed(grid Grid[int]) {
	emptyCell := fmt.Sprintf("%2s", ".")
	cursorCell := fmt.Sprintf("%2s", "@")
	displayGrid(grid, func(cell int) string {
		switch cell {
		case 0:
			return emptyCell
		case greedCursor:
			return cursorCell
		default:
			return fmt.Sprintf("%2d", cell)
		}
	})
}
