package solution_201_250

import "github.com/utmhikari/go-leetcode"

//203. 移除链表元素
//给你一个链表的头节点 head 和一个整数 val ，
//请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

func removeElements203(head *main.ListNode, val int) *main.ListNode {
	dummy := &main.ListNode{Next: head}
	p, prev := head, dummy
	for p != nil {
		if p.Val == val {
			prev.Next = p.Next
			p.Next = nil
			p = prev.Next
		} else {
			p = p.Next
			prev = prev.Next
		}
	}
	h := dummy.Next
	dummy.Next = nil
	return h
}