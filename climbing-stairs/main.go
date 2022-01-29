package main

import "fmt"

func climbStairs(n int) int {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		tmp := b
		b = a + b
		a = tmp
	}
	return a
}

func main() {
	fmt.Println(climbStairs(5))
}
