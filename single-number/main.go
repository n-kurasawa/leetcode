package main

import "fmt"

func singleNumber(nums []int) int {
	var res int
	for _, v := range nums {
		res ^= v
	}
	return res
}

func main() {
	fmt.Println(singleNumber([]int{3, 1, 1}))
}
