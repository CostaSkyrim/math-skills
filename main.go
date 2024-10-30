package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// File reading
	inputFile, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	// Spliting file into lines
	lines := strings.Split(string(inputFile), "\n")

	var numbers []int // This variable holds the numbers

	// Transforming the []byte into a []int
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Printf("Skipping line %q: %v\n", line, err)
			continue
		}
		numbers = append(numbers, num)
	}

	// Printing the results
	fmt.Println("Average:", Average(numbers))
	fmt.Println("Median:", Median(numbers))
	fmt.Println("Variance:", Variance(numbers))
	fmt.Println("Standard Deviation:", StandardDeviation(numbers))
}
