package main

import (
	"fmt"
	"math"
)

func compute(fn func(x, y float64) float64) float64 {
	return fn(3, 4)
}

func main() {

	hypotenus := func(a, b float64) float64 {
		return math.Sqrt(a*a + b*b)
	}

	fmt.Printf("Hypotenus is %v\n", hypotenus(5, 12))
	fmt.Printf("Hypotenus of nested %v\n", compute(hypotenus))
	fmt.Printf("Power is %v\n", compute(math.Pow))
}
