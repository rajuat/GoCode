package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type MyFloat float64

func (mf MyFloat) MyFunction() float64 {
	return float64(mf + 1)
}

func mains() {
	v := Vertex{3, 4}
	y := &v
	z := &Vertex{3, 4}
	fmt.Println(
		v,
		y,
		z,
		*y,
		*z,
	)
	// v := Vertex{X: 3, Y: 4}
	// v.Scale(10)
	// fmt.Println(v.Abs())

	// mf := MyFloat(math.Sqrt2)
	// fmt.Println(mf.MyFunction())
}
