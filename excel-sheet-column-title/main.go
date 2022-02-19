package main

import "fmt"

const base = 'A' - 1

func convertToTitle(columnNumber int) string {
	var b []byte
	for columnNumber > 0 {
		v := columnNumber % 26
		columnNumber /= 26
		if v == 0 {
			v = 26
			columnNumber -= 1
		}
		b = append([]byte{byte(v + base)}, b...)
	}
	return string(b)
}

func main() {
	fmt.Println(convertToTitle(52))
}
