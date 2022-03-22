package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := map[rune]int{}
	for _, v := range s {
		m[v]++
	}
	for _, v := range t {
		m[v]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {}
