package main

import (
	"fmt"
)

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	for i := range row {
		row[i] = 1
	}
	for i := 1; i < rowIndex; i++ {
		for j := i; j > 0; j-- {
			row[j] += row[j-1]
		}
	}
	return row
}

func main() {
	fmt.Println(getRow(1))
	fmt.Println(getRow(4))
}
