package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	mid := middleNode(head)
	reversedSecondHalf := reverse(mid)

	for head != nil && reversedSecondHalf != nil {
		if head.Val != reversedSecondHalf.Val {
			return false
		}
		head = head.Next
		reversedSecondHalf = reversedSecondHalf.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
