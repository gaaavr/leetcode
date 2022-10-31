package main

import "fmt"

// Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := ListNode{
		1,
		&ListNode{
			1,
			&ListNode{
				1,
				&ListNode{
					18,
					&ListNode{
						18,
						nil,
					},
				},
			},
		},
	}
	d := deleteDuplicates(&l)
	fmt.Println(d, d.Next, d.Next.Next)
}
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	if head.Val == head.Next.Val {
		if head.Next.Next == nil {
			head.Next = nil
			return head
		}
		head.Next = head.Next.Next
	}
	if head.Val == head.Next.Val {
		deleteDuplicates(head)
	} else {
		deleteDuplicates(head.Next)
	}
	return head
}
