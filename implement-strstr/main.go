package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	hayRs := []rune(haystack)
	neeRs := []rune(needle)

	for i := 0; i < len(hayRs); i++ {
		if len(hayRs[i:]) < len(neeRs) {
			return -1
		}
		if string(hayRs[i:i+len(neeRs)]) == needle {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaa", "baa"))
	fmt.Println(strStr("", ""))
}
