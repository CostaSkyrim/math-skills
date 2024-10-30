package main

import "math"

func Variance(numbers []int) int {
	n := float64(len(numbers))
	if n == 0 {
		return 0 // Return 0 or another default value if slice is empty
	}

	mean := Average(numbers)

	sumOfSquares := 0.0
	for _, num := range numbers {
		diff := float64(num - mean)
		sumOfSquares += diff * diff
	}

	return int(math.Round(float64(sumOfSquares / n)))
}
