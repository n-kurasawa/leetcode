package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	m := map[*ListNode]bool{}
	for head != nil {
		if m[head] {
			return true
		}
		m[head] = true
		head = head.Next
	}
	return false
}

func main() {
	d := &ListNode{Val: -4, Next: nil}
	c := &ListNode{Val: 0, Next: d}
	b := &ListNode{Val: 2, Next: c}
	a := &ListNode{Val: 3, Next: b}
	fmt.Println(hasCycle(a))
}
