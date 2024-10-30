package main

import (
	"math"
	"sort"
)

func Median(numbers []int) int {
	sort.Ints(numbers) // Sort the slice in ascending order

	n := len(numbers)
	if n == 0 {
		return 0 // Return 0 or another default value if slice is empty
	}

	// If the length is odd, return the middle element
	if n%2 != 0 {
		return numbers[n/2]
	}
	// If the length is even, return the average of the two middle elements (rounded down)
	return int(math.Round((float64(numbers[n/2-1]+numbers[n/2]) / 2)))
}
