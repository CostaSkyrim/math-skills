package main

import (
	"math"
)

func StandardDeviation(numbers []int) int {
	return int(math.Sqrt(float64(Variance(numbers))))
}
