package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func GenLinkedListNode(nums []int) *ListNode {
	var head *ListNode = nil

	if len(nums) == 0 {
		return head
	}

	var tmp *ListNode = nil

	for _, num := range nums {
		node := ListNode{
			Val: num,
		}
		if head == nil {
			head = &node
			tmp = head
		} else if tmp != nil {
			tmp.Next = &node
			tmp = tmp.Next
		}
	}

	return head
}

func (l *ListNode) Equals(ll *ListNode) bool {
	p1, p2 := l, ll

	for {
		if p1 == nil || p2 == nil {
			break
		}

		if p1.Val != p2.Val {
			return false
		}

		p1 = p1.Next
		p2 = p2.Next
	}

	return p1 == nil && p2 == nil
}

func (l *ListNode) ToString() string {
	if l == nil {
		return "nil"
	}

	s := fmt.Sprintf("%d", l.Val)
	tmp := l
	for {
		if tmp.Next == nil {
			return s
		}

		s += fmt.Sprintf("->%d", tmp.Next.Val)
		tmp = tmp.Next
	}
}
