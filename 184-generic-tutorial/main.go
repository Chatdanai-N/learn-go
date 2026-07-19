package main

import "fmt"

func SumInts(m map[string]int64) int64 {

	var summary int64
	for _, v := range m {
		summary += v
	}
	return summary
}

func SumFloat(m map[string]float64) float64 {

	var summary float64
	for _, v := range m {
		summary += v
	}

	return summary
}

type Number interface {
	int64 | float64
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var summary V
	for _, value := range m {
		summary += value
	}
	return summary
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var summary V
	for _, v := range m {
		summary += v
	}
	return summary
}

func main() {

	mpInt := map[string]int64{
		"first":  34,
		"secord": 12,
	}

	mpFloat := map[string]float64{
		"first":  35.98,
		"secord": 26.99,
	}

	fmt.Printf("Non-Generic Sums : %v and %v\n", SumInts(mpInt), SumFloat(mpFloat))
	fmt.Printf("Geneic Sum: %v and %v\n", SumIntsOrFloats[string, int64](mpInt), SumIntsOrFloats[string, float64](mpFloat))
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(mpInt),
		SumIntsOrFloats(mpFloat))
	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(mpInt),
		SumNumbers(mpFloat))
}
