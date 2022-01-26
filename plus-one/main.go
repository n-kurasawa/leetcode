package main

import "fmt"

func plusOne(digits []int) []int {
	var up bool
	for i := len(digits) - 1; i >= 0; i-- {
		plus := digits[i] + 1
		if plus < 10 {
			digits[i] = plus
			return digits
		} else {
			up = true
			digits[i] = 0
		}
	}
	if up {
		digits = append(digits[:1], digits...)
		digits[0] = 1
		return digits
	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{1, 9}))
	fmt.Println(plusOne([]int{9}))
}
