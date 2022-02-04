package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	return 1 + int(math.Max(float64(l), float64(r)))
}

func main() {
	c := &TreeNode{Val: 3, Left: nil, Right: nil}
	b := &TreeNode{Val: 2, Left: c, Right: nil}
	a := &TreeNode{Val: 1, Left: nil, Right: b}
	fmt.Println(maxDepth(a))
}
