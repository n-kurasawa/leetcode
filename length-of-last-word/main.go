package main

import "fmt"

func lengthOfLastWord(s string) int {
	var space bool
	var length int
	for _, v := range []rune(s) {
		if v == ' ' && !space {
			space = true
		} else if v != ' ' && space {
			space = false
			length = 1
		} else if v != ' ' {
			length++
		}
	}
	return length
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
}
