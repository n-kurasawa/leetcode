package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func appendList(first, current, list *ListNode) *ListNode {
	if current == nil {
		first.Next = list
		return list
	}

	current.Next = list
	return current.Next
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	first := &ListNode{}
	var currrent *ListNode
	for list1 != nil || list2 != nil {
		if list1 == nil {
			currrent = appendList(first, currrent, list2)
			list2 = list2.Next
		} else if list2 == nil {
			currrent = appendList(first, currrent, list1)
			list1 = list1.Next
		} else if list1.Val <= list2.Val {
			currrent = appendList(first, currrent, list1)
			list1 = list1.Next
		} else if list1.Val > list2.Val {
			currrent = appendList(first, currrent, list2)
			list2 = list2.Next
		}
	}
	return first.Next
}

func main() {
	fmt.Println(mergeTwoLists(nil, nil))
	list2 := &ListNode{Val: 1, Next: nil}
	fmt.Println(mergeTwoLists(nil, list2))
}
