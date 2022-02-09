package main

import "fmt"

func generate(numRows int) [][]int {
	ret := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		length := i + 1
		ret[i] = make([]int, length)
		for j := 0; j < length; j++ {
			if j == 0 || j == length-1 {
				ret[i][j] = 1
			} else {
				ret[i][j] = ret[i-1][j-1] + ret[i-1][j]
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(generate(5))
}
