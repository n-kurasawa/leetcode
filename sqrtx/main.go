package main

import "fmt"

func mySqrt(x int) int {
	if x == 1 || x == 0 {
		return x
	}
	low := 0
	up := x
	for up-low > 1 {
		mid := (low + up) / 2
		if mid*mid <= x {
			low = mid
		} else {
			up = mid
		}
	}
	return low
}

func main() {
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(1))
	fmt.Println(mySqrt(2147483647))
	fmt.Println(mySqrt(1041080284))
}
