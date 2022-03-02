package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{
		Next: head,
	}
	node := newHead
	for node.Next != nil {
		if node.Next.Val == val {
			node.Next = node.Next.Next
			continue
		}
		node = node.Next
	}
	return newHead.Next
}

func main() {}
