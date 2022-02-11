package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	profit := 0
	minValue := math.MaxInt
	for _, v := range prices {
		minValue = min(v, minValue)
		profit = max(profit, v-minValue)
	}
	return profit
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 2, 1}))
}
