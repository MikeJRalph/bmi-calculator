package main

import (
	"fmt"
	"math"
)

func main() {
	// All units are metric

	// Inputs: weight in kg, height in m
	var weight, height float64

	fmt.Print("Your weight(kg): ")
	fmt.Scan(&weight)

	fmt.Print("Your height(m): ")
	fmt.Scan(&height)

	// Your normal weight range
	minWeight, maxWeight := normalWeightRange(height)
	fmt.Printf("Your normal weight range(min/max): %.1fkg to %.1fkg\n", minWeight, maxWeight)
}

func normalWeightRange(height float64) (normalWeightRangeMin, normalWeightRangeMax float64) {
	// All units are metric
	normalBMIMin := 18.5
	normalBMIMax := 29.4

	normalWeightRangeMin = normalBMIMin * math.Pow(height, 2)
	normalWeightRangeMax = normalBMIMax * math.Pow(height, 2)

	return normalWeightRangeMin, normalWeightRangeMax
}
