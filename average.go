package main

import (
	"math"
)

func Average(numbers []int) int {
	n := float64(len(numbers))
	if n == 0 {
		return 0 // Return 0 or another default value if slice is empty
	}
	sum := 0.0
	for _, num := range numbers {
		sum += float64(num)
	}

	return int(math.Round(sum / n))
}
