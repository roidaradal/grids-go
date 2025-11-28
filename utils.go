package main

import (
	"math/rand/v2"
	"os"
	"os/exec"
	"time"

	"github.com/roidaradal/fn/ds"
)

// Generate <count> random numbers from [0,limit)
func randomNumbers(limit, count int) []int {
	numbers := ds.NewSet[int]()
	for numbers.Len() != count {
		numbers.Add(rand.IntN(limit))
	}
	return numbers.Items()
}

// Run loop
func runLoop[T any](grid Grid[T], displayFn func(Grid[T]), nextFn func(Grid[T]) Grid[T], delayMs int) {
	for {
		clearScreen()
		displayFn(grid)
		if delayMs > 0 {
			time.Sleep(time.Duration(delayMs) * time.Millisecond)
		}
		grid = nextFn(grid)
	}
	// For input:
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// if scanner.Text() == "q" {
	// 	break
	// }
}

// Clear screen
func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
