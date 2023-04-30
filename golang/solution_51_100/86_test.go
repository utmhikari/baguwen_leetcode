package solution_51_100

import "github.com/utmhikari/go-leetcode"

//86. 分隔链表
//给你一个链表的头节点 head 和一个特定值 x ，
//请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
//
//你应当 保留 两个分区中每个节点的初始相对位置。


func partition86(head *main.ListNode, x int) *main.ListNode {
	p := head
	if p == nil || p.Next == nil {
		return head
	}

	geNodesDummy, lNodesDummy := &main.ListNode{}, &main.ListNode{}
	gePtr, lPtr := geNodesDummy, lNodesDummy

	for p != nil {
		next := p.Next
		if p.Val >= x {
			gePtr.Next = p
			gePtr = gePtr.Next
		} else {
			lPtr.Next = p
			lPtr = lPtr.Next
		}
		p.Next = nil
		p = next
	}

	geHead, lHead := geNodesDummy.Next, lNodesDummy.Next
	geNodesDummy.Next = nil
	lNodesDummy.Next = nil
	if geHead == nil {
		return lHead
	}
	if lHead == nil {
		return geHead
	}
	lPtr.Next = geHead
	return lHead
}