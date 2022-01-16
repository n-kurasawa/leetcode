package main

import "fmt"

var (
	roman = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
)

func romanToInt(s string) int {
	int := 0
	for i := len(s) - 1; i >= 0; i-- {
		r := string(s[i])
		if r == "I" || r == "X" || r == "C" {
			if int != 0 && roman[r] < roman[string(s[i+1])] {
				int -= roman[r]
				continue
			}
		}
		int += roman[r]
	}
	return int
}

func main() {
	fmt.Println(3 == romanToInt("III"))
	fmt.Println(58 == romanToInt("LVIII"))
	fmt.Println(1994 == romanToInt("MCMXCIV"))
}
