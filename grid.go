package main

import (
	"fmt"
	"strings"

	"github.com/roidaradal/fn/ds"
)

type Grid[T any] [][]T

// Create new blank grid
func newGrid[T any](numRows, numCols int) Grid[T] {
	grid := make(Grid[T], numRows)
	for i := range numRows {
		grid[i] = make([]T, numCols)
	}
	return grid
}

// Return numRows, numCols of grid
func (grid Grid[T]) Shape() (int, int) {
	return len(grid), len(grid[0])
}

// Convert index to Coords
func indexToCoords(idx, width int) ds.Coords {
	row, col := idx/width, idx%width
	return ds.Coords{row, col}
}

// Count the neighbors of grid cell at <coords> using the ok function
func countNeighbors[T any](grid Grid[T], coords ds.Coords, ok func(T) bool) int {
	numRows, numCols := grid.Shape()
	y, x := coords.Tuple()
	count := 0
	deltas := []int{-1, 0, 1}
	for _, dy := range deltas {
		for _, dx := range deltas {
			if dy == 0 && dx == 0 {
				continue
			}
			ny, nx := y+dy, x+dx
			if notInsideBounds(ny, nx, numRows, numCols) {
				continue
			}
			if ok(grid[ny][nx]) {
				count += 1
			}
		}
	}
	return count
}

// Check if (y,x) coords is outside of the (numRows, numCols) grid
func notInsideBounds(y, x, numRows, numCols int) bool {
	return y < 0 || x < 0 || y >= numRows || x >= numCols
}

// Display grid using the toString function
func displayGrid[T any](grid Grid[T], toString func(T) string) {
	for _, line := range grid {
		out := make([]string, len(line))
		for i, cell := range line {
			out[i] = toString(cell)
		}
		fmt.Println(strings.Join(out, ""))
	}
}
