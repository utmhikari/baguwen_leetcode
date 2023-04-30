package solution_101_150

import "github.com/utmhikari/go-leetcode"

//148. 排序链表
//给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。



func merge148(left *main.ListNode, right *main.ListNode) *main.ListNode {
	dummy := &main.ListNode{
		Next: left,
	}

	p0, p1, p2 := dummy, left, right

	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			p0.Next = p1
			p1 = p1.Next
		} else {
			p0.Next = p2
			p2 = p2.Next
		}

		p0 = p0.Next
	}

	if p1 != nil {
		p0.Next = p1
	} else {
		p0.Next = p2
	}

	h := dummy.Next
	dummy.Next = nil
	return h
}


func sortListWithTail148(head *main.ListNode, tail *main.ListNode) *main.ListNode {
	if nil == head {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	p1, p2 := head, head
	for p2 != tail {
		p1 = p1.Next
		p2 = p2.Next
		if p2 != tail {
			p2 = p2.Next
		}
	}

	left, right := sortListWithTail148(head, p1), sortListWithTail148(p1, tail)
	return merge148(left, right)
}



func sortList148(head *main.ListNode) *main.ListNode {
	return sortListWithTail148(head, nil)
}
