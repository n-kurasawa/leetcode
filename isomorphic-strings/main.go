package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	sm := map[byte]byte{}
	tm := map[byte]byte{}
	for i := 0; i < len(s); i++ {
		if mt, ok := sm[s[i]]; ok {
			if mt != t[i] {
				fmt.Println("t", i)
				return false
			}
			continue
		}
		if ms, ok := tm[t[i]]; ok {
			if ms != s[i] {
				fmt.Println("s", i)
				return false
			}
			continue
		}

		sm[s[i]] = t[i]
		tm[t[i]] = s[i]
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("paper", "title"))
}
