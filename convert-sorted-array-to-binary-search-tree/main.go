package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	center := len(nums) / 2
	return &TreeNode{
		Val:   nums[center],
		Left:  sortedArrayToBST(nums[:center]),
		Right: sortedArrayToBST(nums[center+1:]),
	}
}

func main() {
	sortedArrayToBST([]int{-10, -3, 0, 5, 9})
}
