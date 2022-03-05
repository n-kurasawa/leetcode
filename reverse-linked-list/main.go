package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Stack struct {
	values []*ListNode
}

func (s *Stack) push(n *ListNode) {
	s.values = append(s.values, n)
}

func (s *Stack) pop() *ListNode {
	if len(s.values) == 0 {
		return nil
	}
	n := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return n
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	s := &Stack{}
	for head != nil {
		next := head.Next
		s.push(head)
		head = next
	}
	h := s.pop()
	n := h
	for n != nil {
		next := s.pop()
		n.Next = next
		n = next
	}
	return h
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func main() {
	b := &ListNode{Val: 2, Next: nil}
	a := &ListNode{Val: 1, Next: b}
	printList(reverseList(a))
}
