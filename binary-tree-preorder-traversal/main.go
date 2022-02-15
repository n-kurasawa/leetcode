package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	return traversal(root, []int{})
}

func traversal(node *TreeNode, acc []int) []int {
	if node == nil {
		return acc
	}
	acc = append(acc, node.Val)
	acc = traversal(node.Left, acc)
	acc = traversal(node.Right, acc)
	return acc
}

func main() {}
