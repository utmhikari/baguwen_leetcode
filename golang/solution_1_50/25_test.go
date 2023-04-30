package solution_1_50

import "github.com/utmhikari/go-leetcode"

//25. K 个一组翻转链表
//给你一个链表，每k个节点一组进行翻转，请你返回翻转后的链表。
//
//k是一个正整数，它的值小于或等于链表的长度。
//
//如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。
//
//进阶：
//
//你可以设计一个只使用常数额外空间的算法来解决此问题吗？
//你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。


func reverse25(head *main.ListNode) {
	var rev *main.ListNode = nil
	var tmp *main.ListNode
	p := head
	for p != nil {
		tmp = p.Next
		p.Next = rev
		rev = p
		p = tmp
	}
}



func reverseKGroup25(head *main.ListNode, k int) *main.ListNode {
	if head == nil {
		return nil
	}
	if k == 1 {
		return head
	}

	dummy := &main.ListNode{Next: head}
	pivot := dummy
	for {
		p := pivot
		cnt := 0
		for p.Next != nil && cnt < k {
			p = p.Next
			cnt++
		}
		if cnt < k {
			break
		}

		h, t := pivot.Next, p
		tailNext := t.Next
		pivot.Next = nil
		t.Next = nil
		reverse25(h)
		pivot.Next = t
		h.Next = tailNext
		pivot = h
	}

	ans := dummy.Next
	dummy.Next = nil
	return ans
}