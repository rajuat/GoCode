package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"strings"
	"time"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Input type is %T value is %v\n", ToBe, ToBe)
	fmt.Printf("Input type is %T value is %v\n", MaxInt, MaxInt)
	fmt.Printf("Input type is %T value is %v\n", z, z)

	a, b := split(55)
	fmt.Printf("Splited values are %d %d \n", a, b)

	x, y := swap("hello", "world")
	fmt.Println(x, y)

	fmt.Printf("hello")
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

	fmt.Println("The time is ", time.Now())

	nums := []int{0, 2, 3, 4, 8, 6, 1}
	fmt.Printf("Missing number is: %d \n", findMissingInteger(nums))

	fmt.Printf("Power is %v\n", pow(3, 3, 10))

	UnderstandingPointers()

	TicTacToe()

	Pow()

	RangeContinued()

	MapExample()

	var m map[string]uint = WordCount("This is my name and is good")
	fmt.Println(m)
}

func WordCount(s string) map[string]uint {

	var words []string = strings.Fields(s)
	var m = make(map[string]uint)
	for i := 0; i < len(words); i++ {
		v, ok := m[words[i]]
		if ok {
			m[words[i]] = v + 1
		} else {
			m[words[i]] = 1
		}
	}

	return m
}

type Vertex struct {
	Lat, Lon float64
}

func MapExample() {
	m := make(map[string]Vertex)
	m["Delhi"] = Vertex{
		Lat: 100,
		Lon: 200,
	}
	fmt.Println(m["Delhi"])

	literal := map[string]Vertex{
		"Bell Labs": {1000, 2000},
		"Google":    {3000, 4000},
	}

	fmt.Println(literal)
}

func RangeContinued() {
	pow := make([]int, 10)
	for index := range pow {
		pow[index] = 1 << uint(index)
	}

	for _, value := range pow {
		defer fmt.Printf("Multiplies %d\n", value)
	}
}

func Pow() {
	pow := []int{1, 2, 4, 8, 16, 32, 64}
	for i, v := range pow {
		fmt.Printf("2*%d = %d\n", i, v)
	}

}

func TicTacToe() {
	board := [][]string{
		[]string{"X", "-", "Y"},
		[]string{"Y", "X", "-"},
		[]string{"-", "X", "Y"},
	}

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func UnderstandingPointers() {
	i := 42
	p := &i
	fmt.Printf("Value of p %v\n", *p)
	*p = 21
	fmt.Printf("Value of i %d\n", i)

}

func pow(x, n, lim float64) float64 {
	if y := math.Pow(x, n); y <= lim {
		return y
	} else {
		fmt.Printf("%g >= %g\n", y, lim)
	}
	return lim
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func findMissingInteger(nums []int) int {
	n := 1
	i := 0

	for ; i < len(nums); i++ {
		if nums[i] == n {
			n++
			i = -1
		}
	}

	return n
}

func swap(x, y string) (string, string) {
	return y, x
}
