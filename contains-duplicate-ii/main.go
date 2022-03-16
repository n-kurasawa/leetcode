package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	for i, v := range nums {
		index, ok := m[v]
		if ok && abs(i-index) <= k {
			return true
		}
		m[v] = i
	}
	return false
}

func abs(n int) int {
	if n > 0 {
		return n
	} else {
		return -n
	}
}

func main() {}
