package solution_51_100

import (
	"fmt"
	"github.com/utmhikari/go-leetcode"
)

//92. 反转链表 II
//给你单链表的头指针 head 和两个整数left 和 right ，
//其中left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。



func reverseBetween(head *main.ListNode, left int, right int) *main.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var leftPrev *main.ListNode
	var rightNext *main.ListNode

	dummy := &main.ListNode{}
	dummy.Next = head

	p := dummy
	cnt := 1
	for p != nil {
		if p.Next != nil && cnt == left {
			leftPrev = p
		}
		if cnt > right {
			rightNext = p.Next
			p.Next = nil
			break
		}
		p = p.Next
		cnt++
	}

	// invalid input
	if leftPrev == nil || right < left {
		fmt.Printf("invalid input: %d, %d\n", left, right)
		return head
	}

	// reverse
	p = leftPrev.Next
	tail := rightNext
	var tmp *main.ListNode
	leftPrev.Next = nil
	for {
		tmp = p.Next
		p.Next = tail
		tail = p
		p = tmp
		if tmp == nil {
			break
		}
	}
	leftPrev.Next = tail

	// get head
	newHead := dummy.Next
	dummy.Next = nil
	return newHead
}
