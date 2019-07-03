package main

import (
	"fmt"
	"time"
)

func mains() {
	fmt.Printf("sqrt is %v\n", sqrt(2))
	fmt.Printf("sqrt is %v\n", sqrt(3))
	fmt.Printf("sqrt is %v\n", sqrt(4))
	fmt.Printf("sqrt is %v\n", sqrt(9))

	WhenIsSaturday()
}

func sqrt(x float64) float64 {
	var z float64 = 1
	for z*z < x {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func WhenIsSaturday() {
	today := time.Now().Weekday()
	switch time.Saturday {
	case today:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
}
