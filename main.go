package main

import "fmt"

const (
	MINES = "mines"
	GREED = "greed"
	_2048 = "2048"
)

func main() {
	option := _2048
	switch option {
	case MINES:
		grid := newMines(9, 9, 10)
		displayMines(grid)
	case GREED:
		grid := newGreed()
		displayGreed(grid)
	case _2048:
		grid := new2048()
		display2048(grid)
	default:
		fmt.Println("Unknown option:", option)
	}
}
