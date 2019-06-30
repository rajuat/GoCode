package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("hello")

	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

	fmt.Println("The time is ", time.Now())

	nums := []int{0, 2, 3, 4, 8, 6, 1}
	fmt.Printf("Missing number is: %d", findMissingInteger(nums))
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
