package main

import (
	"math/rand/v2"
	"slices"

	"github.com/roidaradal/fn/ds"
)

// TODO: replace with check.IsEqual
func checkIsEqual[T comparable](goalValue T) func(T) bool {
	return func(value T) bool {
		return value == goalValue
	}
}

// TODO: replace with list.Repeated
func listRepeated[T any](value T, count int) []T {
	return slices.Repeat([]T{value}, count)
}

// Generate <count> random numbers from [0,limit)
func randomNumbers(limit, count int) []int {
	numbers := ds.NewSet[int]()
	for numbers.Len() != count {
		numbers.Add(rand.IntN(limit))
	}
	return numbers.Items()
}
