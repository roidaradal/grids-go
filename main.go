package main

import "fmt"

const (
	_2048  = "2048"
	CONWAY = "conway"
	GREED  = "greed"
	MINES  = "mines"
)

func main() {
	option := CONWAY
	switch option {
	case _2048:
		grid := new2048()
		display2048(grid)
	case CONWAY:
		grid := newRandomConway(50, 250, 500)
		runLoop(grid, displayConway, nextConway, 100)
	case GREED:
		grid := newGreed()
		displayGreed(grid)
	case MINES:
		grid := newMines(9, 9, 10)
		displayMines(grid)
	default:
		fmt.Println("Unknown option:", option)
	}
}
