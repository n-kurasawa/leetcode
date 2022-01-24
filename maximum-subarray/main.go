package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	max := math.MinInt64
	localMax := 0
	for i := 0; i < len(nums); i++ {
		localMax = int(math.Max(float64(nums[i]), float64(nums[i]+localMax)))
		if localMax > max {
			max = localMax
		}
	}
	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{-1}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
	fmt.Println(maxSubArray([]int{10, -1, -1}))
	fmt.Println(maxSubArray([]int{-1, 10, -1}))
	fmt.Println(maxSubArray([]int{-1, -1, 10}))
}
