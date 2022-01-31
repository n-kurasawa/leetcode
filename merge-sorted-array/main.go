package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := len(nums1)-1, m-1, n-1
	for j >= 0 && k >= 0 {
		if nums1[j] > nums2[k] {
			nums1[i] = nums1[j]
			j--
		} else {
			nums1[i] = nums2[k]
			k--
		}
		i--
	}
	for k >= 0 {
		nums1[i] = nums2[k]
		k--
		i--
	}
}

func main() {
	sl := []int{1, 2, 3, 0, 0, 0}
	merge(sl, 3, []int{2, 5, 6}, 3)
	fmt.Println(sl)

	sl2 := []int{1}
	merge(sl2, 1, []int{}, 0)
	fmt.Println(sl2)

	sl3 := []int{0}
	merge(sl3, 0, []int{1}, 1)
	fmt.Println(sl3)

	sl4 := []int{4, 5, 6, 0, 0, 0}
	merge(sl4, 3, []int{1, 2, 3}, 3)
	fmt.Println(sl4)
}
