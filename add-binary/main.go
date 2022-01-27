package main

import (
	"fmt"
	"strconv"
)

func runeToInt(r rune) int {
	i, _ := strconv.Atoi(string(r))
	return i
}

func intToRune(i int) rune {
	return []rune(strconv.Itoa(i))[0]
}

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		diff := len(b) - len(a)
		for i := 0; i < diff; i++ {
			a = "0" + a
		}
	} else {
		diff := len(a) - len(b)
		for i := 0; i < diff; i++ {
			b = "0" + b
		}
	}

	up := 0
	ars := []rune(a)
	brs := []rune(b)
	for i := len(ars) - 1; i >= 0; i-- {
		ai := runeToInt(ars[i])
		bi := runeToInt(brs[i])
		sum := ai + bi + up
		if sum < 2 {
			ars[i] = intToRune(sum)
			up = 0
		} else {
			up = 1
			ars[i] = intToRune(sum - 2)
		}
	}

	if up > 0 {
		ars = append(ars[:1], ars...)
		ars[0] = intToRune(1)
		return string(ars)
	}
	return string(ars)
}

func main() {
	fmt.Println(addBinary("11", "11"))
	fmt.Println(addBinary("1010", "1011"))
}
