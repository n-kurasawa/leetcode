package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return checkSymmetric(root.Left, root.Right)
}

func checkSymmetric(first, seccond *TreeNode) bool {
	if first == nil && seccond == nil {
		return true
	}
	if first != nil && seccond == nil {
		return false
	}
	if first == nil && seccond != nil {
		return false
	}
	if first.Val != seccond.Val {
		return false
	}
	return checkSymmetric(first.Left, seccond.Right) && checkSymmetric(first.Right, seccond.Left)
}

func main() {}
