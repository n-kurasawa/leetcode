package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var acc []int
	return travers(root, acc)

}

func travers(node *TreeNode, acc []int) []int {
	if node != nil {
		acc = travers(node.Left, acc)
		acc = append(acc, node.Val)
		acc = travers(node.Right, acc)
	}
	return acc
}

func main() {
	c := &TreeNode{Val: 3, Left: nil, Right: nil}
	b := &TreeNode{Val: 2, Left: c, Right: nil}
	a := &TreeNode{Val: 1, Left: nil, Right: b}
	fmt.Println(inorderTraversal(a))
}
