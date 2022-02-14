package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if slow == nil || fast == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func main() {
	d := &ListNode{Val: -4, Next: nil}
	c := &ListNode{Val: 0, Next: d}
	b := &ListNode{Val: 2, Next: c}
	a := &ListNode{Val: 3, Next: b}
	fmt.Println(hasCycle(a))
}
