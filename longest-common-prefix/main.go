package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for _, v := range strs[1:] {
		tmp := ""
		for i := 0; i < len(v) && i < len(prefix); i++ {
			if v[i] == prefix[i] {
				tmp += string(v[i])
			} else {
				break
			}
		}
		prefix = tmp
	}
	return prefix
}

func main() {
	fmt.Println("fl" == longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println("" == longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println("a" == longestCommonPrefix([]string{"ab", "a"}))
}
