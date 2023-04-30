package solution_51_100

import "github.com/utmhikari/go-leetcode"

//61. 旋转链表
//给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight61(head *main.ListNode, k int) *main.ListNode {
	if head == nil || head.Next == nil {
		// no need to rotate
		return head
	}

	// get length
	var tail *main.ListNode = head
	p, l := head, 0
	for p != nil {
		l++
		if p.Next == nil {
			tail = p
		}
		p = p.Next
	}

	// move l == move 0
	k %= l
	if k == 0 {
		// no need to rotate
		return head
	}

	// get last k node
	h := head
	p = h
	for i := 0; i < l - k - 1; i++ {
		p = p.Next
	}
	tail.Next = h
	ans := p.Next
	p.Next = nil
	return ans
}