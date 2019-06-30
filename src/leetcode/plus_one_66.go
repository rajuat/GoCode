package main

// func main() {
// 	digits := []int{9, 9, 9, 9, 9}
// 	fmt.Println(plusOne(digits))
// }

func plusOne(digits []int) []int {
	length := len(digits)
	carry := 0
	toBeAdded := 1
	for i := length - 1; i > -1; i-- {
		digit := digits[i] + carry + toBeAdded
		toBeAdded = 0
		if digit < 10 {
			digits[i] = digit
			return digits
		} else {
			singleDigit := digit % 10
			carry = digit / 10
			digits[i] = singleDigit
		}
	}
	if carry > 0 {
		result := make([]int, length+1)
		result[0] = carry
		for j := 0; j < length; j++ {
			result[j+1] = digits[j]
		}
		return result
	}

	return nil
}
