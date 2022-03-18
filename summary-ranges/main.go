package main

import "strconv"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	ans, l, r := []string{}, 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 {
			r = i
		} else {
			if l == r {
				ans = append(ans, strconv.Itoa(nums[l]))
			} else {
				ans = append(ans, strconv.Itoa(nums[l])+"->"+strconv.Itoa(nums[r]))
			}
			l, r = i, i
		}
	}
	if l == r {
		return append(ans, strconv.Itoa(nums[l]))
	}
	return append(ans, strconv.Itoa(nums[l])+"->"+strconv.Itoa(nums[r]))
}

func main() {}
