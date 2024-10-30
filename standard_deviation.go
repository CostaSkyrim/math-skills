package main

import (
	"math"
)

func StandardDeviation(numbers []int) int {
	return int(math.Round(math.Sqrt(float64(Variance(numbers)))))
}
