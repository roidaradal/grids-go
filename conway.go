package main

import (
	"github.com/roidaradal/fn/ds"
	"github.com/roidaradal/fn/lang"
)

// Create new random Conway's Game of Life grid
func newRandomConway(numRows, numCols, numCells int) Grid[bool] {
	grid := newGrid[bool](numRows, numCols)
	for _, idx := range randomNumbers(numRows*numCols, numCells) {
		row, col := indexToCoords(idx, numCols).Tuple()
		grid[row][col] = true
	}
	return grid
}

// Compute the next Conway's Game of Life grid
func nextConway(grid Grid[bool]) Grid[bool] {
	numRows, numCols := grid.Shape()
	grid2 := make(Grid[bool], numRows)
	isCellAlive := func(isAlive bool) bool {
		return isAlive
	}
	for row, line := range grid {
		line2 := make([]bool, numCols)
		for col, currAlive := range line {
			count := countNeighbors(grid, ds.Coords{row, col}, isCellAlive)
			nextAlive := false
			if currAlive && count < 2 {
				nextAlive = false
			} else if currAlive && (count == 2 || count == 3) {
				nextAlive = true
			} else if currAlive && count > 3 {
				nextAlive = false
			} else if !currAlive && count == 3 {
				nextAlive = true
			}
			line2[col] = nextAlive
		}
		grid2[row] = line2
	}
	return grid2
}

// Display Conway's Game of Life grid
func displayConway(grid Grid[bool]) {
	displayGrid(grid, func(isAlive bool) string {
		return lang.Ternary(isAlive, "\u25A0", " ")
	})
}
