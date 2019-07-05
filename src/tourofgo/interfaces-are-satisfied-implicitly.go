package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()

	var a I
	var t = T{"world"}
	a = t
	a.M()

}
