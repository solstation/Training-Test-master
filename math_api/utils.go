package main

import (
	"math"
	"sort"
)

// min return the min number of an array
func Min(numbers []float64) float64 {
	min := math.Inf(1)

	for _, val := range numbers {
		if val < min {
			min = val
		}
	}

	return min
}

// min return the max number of an array
func Max(numbers []float64) float64 {
	max := math.Inf(-1)

	for _, val := range numbers {
		if val > max {
			max = val
		}
	}

	return max
}

// Average return the average of an array
func Average(numbers []float64) float64 {
	var avg float64 = 0

	for _, val := range numbers {
		avg += val
	}

	return avg / float64(len(numbers))
}

// Average return the median of an array
func Median(numbers []float64) float64 {
	var median int = int(math.Ceil(float64(len(numbers) / 2)))
	var med float64 = numbers[median]

	sort.Float64s(numbers)

	if len(numbers)%2 == 0 {
		med = (med + numbers[median-1]) / 2.0
	}

	return med
}

// Percentile return the amount'th percentile of the list
func Percentile(numbers Numbers) float64 {
	perc := (numbers.Amount / 100.0) * float64(len(numbers.Numbers))

	return numbers.Numbers[int(math.Ceil(perc))-1]
}
