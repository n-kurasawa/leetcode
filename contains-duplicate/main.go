package main

func containsDuplicate(nums []int) bool {
	m := map[int]bool{}
	for _, v := range nums {
		if m[v] {
			return true
		}
		m[v] = true
	}
	return false
}

func main() {

}
