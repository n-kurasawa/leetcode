package main

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}
	if n < 3 {
		return true
	}

	num := 2
	for i := 1; i < n; i++ {
		num *= 2
		if num == n {
			return true
		}
	}
	return false
}

func main() {}
