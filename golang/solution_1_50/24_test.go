package solution_1_50

import "github.com/utmhikari/go-leetcode"

//24. 两两交换链表中的节点
//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。




func swapPairs24(head *main.ListNode) *main.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &main.ListNode{Next: head}
	pivot := dummy
	var first, second *main.ListNode
	for {
		first = pivot.Next
		if first == nil {
			break
		}
		second = first.Next
		if second == nil {
			break
		}
		first.Next = second.Next
		pivot.Next = second
		second.Next = first
		pivot = first
	}

	ans := dummy.Next
	dummy.Next = nil
	return ans
}
