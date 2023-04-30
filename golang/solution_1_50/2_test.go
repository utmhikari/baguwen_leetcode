package solution_1_50

import "github.com/utmhikari/go-leetcode"

//2. 两数相加
//给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
//
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//
//你可以假设除了数字 0 之外，这两个数都不会以 0开头。


func addTwoNumbers2(l1 *main.ListNode, l2 *main.ListNode) *main.ListNode {
	if nil == l1 {
		return l2
	}
	if nil == l2 {
		return l1
	}

	p1, p2 := l1, l2
	dummy := &main.ListNode{}
	node := dummy
	inc := 0

	for p1 != nil || p2 != nil {
		sm := inc
		if p1 != nil {
			sm += p1.Val
		}
		if p2 != nil {
			sm += p2.Val
		}
		if sm >= 10 {
			sm -= 10
			inc = 1
		} else {
			inc = 0
		}
		node.Next = &main.ListNode{Val: sm}
		node = node.Next
		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
	}

	if inc == 1 {
		node.Next = &main.ListNode{Val: 1}
	}

	return dummy.Next
}
