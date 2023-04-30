package solution_101_150

import "github.com/utmhikari/go-leetcode"

//143. 重排链表
//给定一个单链表L：L0→L1→…→Ln-1→Ln ，
//将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
//
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。


func reverseList143(head *main.ListNode) *main.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var dummy *main.ListNode
	p0 := head
	var p1 *main.ListNode
	for {
		p1 = p0.Next
		p0.Next = dummy
		dummy = p0
		p0 = p1

		if p1 == nil {
			break
		}
	}

	return dummy
}


func reorderList143(head *main.ListNode)  {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	dummy := &main.ListNode{
		Next: head,
	}

	p1, p2 := dummy, dummy

	for {
		p1 = p1.Next
		p2 = p2.Next
		if p2 == nil {
			break
		}
		p2 = p2.Next
		if p2 == nil {
			break
		}
	}

	p3 := p1.Next
	p1.Next = nil
	p3 = reverseList143(p3)

	p1 = dummy.Next
	var p4 *main.ListNode
	for {
		p2 = p1.Next
		p4 = p3.Next
		p1.Next = p3
		p3.Next = p2
		p3 = p4
		p1 = p2
		if p3 == nil || p1 == nil {
			break
		}
	}

	dummy.Next = nil
}