package main

import "fmt"

const (
	MINES = "mines"
)

func main() {
	option := MINES
	switch option {
	case MINES:
		grid := newMines(9, 9, 10)
		displayMines(grid)
	default:
		fmt.Println("Unknown option:", option)
	}
}
