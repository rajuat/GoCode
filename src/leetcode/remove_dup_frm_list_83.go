package main

import "fmt"

func mains() {
	n5 := &ListNode{5, nil}
	n4 := &ListNode{4, n5}
	n3 := &ListNode{4, n4}
	n2 := &ListNode{2, n3}
	n1 := &ListNode{1, n2}
	n := deleteDuplicates(n1)
	for r := n; r != nil; r = r.Next {

		fmt.Println(r.Val)
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	for runner := head; runner != nil && runner.Next != nil; {
		if runner.Val == runner.Next.Val {
			runner.Next = runner.Next.Next
		} else {
			runner = runner.Next
		}
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
