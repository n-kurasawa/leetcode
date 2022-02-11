package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	profit := 0
	max := 0
	for i := len(prices) - 1; i > 0; i-- {
		if prices[i] <= max {
			continue
		}
		max = prices[i]
		for j := i - 1; j >= 0; j-- {
			if profit < prices[i]-prices[j] {
				profit = prices[i] - prices[j]
			}
		}
	}
	return profit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 2, 1}))
}
