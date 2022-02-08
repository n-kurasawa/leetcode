package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return hasSum(root, targetSum, 0)
}

func hasSum(node *TreeNode, targetSum, sum int) bool {
	s := sum + node.Val
	if node.Left == nil && node.Right == nil {
		return targetSum == s
	}
	if node.Right == nil {
		return hasSum(node.Left, targetSum, s)
	}
	if node.Left == nil {
		return hasSum(node.Right, targetSum, s)
	}
	return hasSum(node.Left, targetSum, s) || hasSum(node.Right, targetSum, s)
}

func main() {}
