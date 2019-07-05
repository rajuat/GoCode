package main

import (
	"fmt"
	"math"
)

func compute(fn func(x, y float64) float64) float64 {
	return fn(3, 4)
}

func mains() {

	hypotenus := func(a, b float64) float64 {
		return math.Sqrt(a*a + b*b)
	}

	fmt.Printf("Hypotenus is %v\n", hypotenus(5, 12))
	fmt.Printf("Hypotenus of nested %v\n", compute(hypotenus))
	fmt.Printf("Power is %v\n", compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	first := 0
	second := 1
	return func() int {
		sum := first + second
		first = second
		second = sum
		return sum
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
