package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	c := &TreeNode{Val: 3, Left: nil, Right: nil}
	b := &TreeNode{Val: 2, Left: nil, Right: nil}
	a := &TreeNode{Val: 1, Left: b, Right: c}

	fmt.Println(isSameTree(a, a))

	e := &TreeNode{Val: 2, Left: nil, Right: nil}
	d := &TreeNode{Val: 1, Left: e, Right: nil}

	g := &TreeNode{Val: 2, Left: nil, Right: nil}
	f := &TreeNode{Val: 1, Left: nil, Right: g}

	fmt.Println(isSameTree(d, f))
}
