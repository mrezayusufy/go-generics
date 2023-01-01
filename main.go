package main

import "fmt"

type Number interface {
	int64 | float64
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s

}

// SumIntsOrFloats sums the values of map m. it supports both int64 and float64
// as types for map values
func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	// initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// initialize a map for the float values
	floats := map[string]float64{
		"first":  3.14,
		"second": 34.56,
	}

	fmt.Print("\n----Non-Generics----\n")
	fmt.Printf("Non-Generics sumsInts: %v\n", SumInts(ints))
	fmt.Printf("Non-Generics sumFloats: %v", SumFloats(floats))

	fmt.Print("\n----Generics----\n")
	fmt.Printf("Generics Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))
	fmt.Printf("Generics Sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))
}
