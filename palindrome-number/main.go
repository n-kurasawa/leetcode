package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	i := 0
	for j := len(str) - 1; j >= 0; j-- {
		if str[i] != str[j] {
			return false
		}
		i++
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}
