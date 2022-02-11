package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	l := strings.ToLower(s)
	r := []rune(l)
	for i := 0; i < len(r); i++ {
		if !isAlphanumeric(int(r[i])) {
			r = append(r[:i], r[i+1:]...)
			i--
		}
	}
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			return false
		}
	}
	return true
}

func isAlphanumeric(c int) bool {
	if '0' <= c && c <= '9' {
		return true
	}
	if 'a' <= c && c <= 'z' {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPalindrome(" "))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("aa"))
	fmt.Println(isPalindrome("aaa"))
}
