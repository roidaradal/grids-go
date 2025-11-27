package main

import (
	"math/rand/v2"

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
