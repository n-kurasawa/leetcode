package main

import "fmt"

func isHappy(n int) bool {
	seen := map[int]bool{}
	for n != 1 && !seen[n] {
		seen[n] = true
		n = nextNum(n)
	}
	return n == 1
}

func nextNum(n int) int {
	var sum int
	for n > 0 {
		d := n % 10
		n = n / 10
		sum += d * d
	}
	return sum
}

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
}
