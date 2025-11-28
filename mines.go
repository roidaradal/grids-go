package main

import (
	"fmt"

	"github.com/roidaradal/fn/check"
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/lang"
)

// Easy		9x9		10
// Medium	16x16	40
// Hard 	30x16	99

const bomb int = -1

// Create new Minesweeper grid
func newMines(numRows, numCols, numMines int) Grid[int] {
	// Make blank int grid
	grid := newGrid[int](numRows, numCols)
	// Randomly place bombs
	for _, idx := range randomNumbers(numRows*numCols, numMines) {
		row, col := indexToCoords(idx, numCols).Tuple()
		grid[row][col] = bomb
	}
	isBomb := check.IsEqual(bomb)
	// Count neighbor mines
	for row, line := range grid {
		for col, cell := range line {
			if cell == bomb {
				continue
			}
			grid[row][col] = countNeighbors(grid, ds.Coords{row, col}, isBomb)
		}
	}
	return grid
}

// Display Minesweeper grid
func displayMines(grid Grid[int]) {
	bombCell := fmt.Sprintf("%3s", "X")
	displayGrid(grid, func(cell int) string {
		normalCell := fmt.Sprintf("%3d", cell)
		return lang.Ternary(cell == bomb, bombCell, normalCell)
	})
}
