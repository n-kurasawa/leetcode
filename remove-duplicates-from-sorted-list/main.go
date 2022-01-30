package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	node := head
	for node != nil {
		if node.Next == nil {
			break
		}
		if node.Val == node.Next.Val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return head
}

func printNodes(head *ListNode) {
	node := head
	for node != nil {
		fmt.Printf("%d, ", node.Val)
		node = node.Next
	}
	fmt.Print("\n")
}

func main() {
	f := &ListNode{Val: 3, Next: nil}
	e := &ListNode{Val: 3, Next: f}
	d := &ListNode{Val: 3, Next: e}
	c := &ListNode{Val: 2, Next: d}
	b := &ListNode{Val: 1, Next: c}
	a := &ListNode{Val: 1, Next: b}
	printNodes(a)
	printNodes(deleteDuplicates(a))
}
