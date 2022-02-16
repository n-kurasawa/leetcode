package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	return traversal(root, []int{})
}

func traversal(node *TreeNode, acc []int) []int {
	if node == nil {
		return acc
	}
	acc = traversal(node.Left, acc)
	acc = traversal(node.Right, acc)
	acc = append(acc, node.Val)
	return acc
}

func main() {}
