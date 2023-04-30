package solution_201_250

import "github.com/utmhikari/go-leetcode"

//206. 反转链表


func reverseList206(head *main.ListNode) *main.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var h, nxt *main.ListNode
	p := head
	for p != nil {
		nxt = p.Next
		p.Next = h
		h = p
		p = nxt
	}
	return h
}
