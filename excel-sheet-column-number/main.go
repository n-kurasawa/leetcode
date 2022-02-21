package main

import (
	"fmt"
)

func titleToNumber(s string) int {
	col := 0
	for _, r := range s {
		cur := int(r - 'A' + 1)
		fmt.Println(cur)
		col = col*26 + cur
	}
	return col
}

func main() {
	fmt.Println(titleToNumber("ZY"))
}
