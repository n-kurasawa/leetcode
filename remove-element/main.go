package main

import "fmt"

func remove(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = remove(nums, i)
			i--
		}
	}
	return len(nums)
}

func main() {
	nums := []int{3, 2, 2, 3}
	fmt.Println(removeElement(nums, 3))
	fmt.Println(nums)
}
