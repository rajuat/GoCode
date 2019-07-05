package main

import "fmt"

type MinStack struct {
	Array        []int
	MinIndex     int
	NextMinIndex int
	TopIndex     int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var A []int
	minStack := MinStack{A, 0, 0, -1}
	return minStack
}

func (this *MinStack) Push(x int) {
	this.TopIndex++
	this.Array = append(this.Array, x)
	if x < this.Array[this.MinIndex] {
		this.NextMinIndex = this.MinIndex
		this.MinIndex = this.TopIndex
	}
}

func (this *MinStack) Pop() {
	if this.TopIndex > -1 {
		if this.TopIndex == this.MinIndex {
			this.MinIndex = this.NextMinIndex
		}
		this.TopIndex--
	}
}

func (this *MinStack) Top() int {
	if this.TopIndex > -1 {
		return this.Array[this.TopIndex]
	} else {
		return 0
	}
}

func (this *MinStack) GetMin() int {
	MinValue := this.Array[this.MinIndex]
	return MinValue
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	param_1 := obj.GetMin()
	obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.GetMin()

	fmt.Println(
		param_1,
		param_3,
		param_4,
	)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
