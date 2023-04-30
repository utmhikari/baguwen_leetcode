package solution_1_50

import "github.com/utmhikari/go-leetcode"

//23. 合并K个升序链表
//给你一个链表数组，每个链表都已经按升序排列。
//
//请你将所有链表合并到一个升序链表中，返回合并后的链表。



func merge23(lists []*main.ListNode, left int, right int) *main.ListNode {
	if right < left {
		return nil
	}
	if right == left {
		return lists[left]
	}

	mid := left + (right - left) / 2
	leftNodes := merge23(lists, left, mid)
	rightNodes := merge23(lists, mid + 1, right)

	dummy := &main.ListNode{}
	tmp := dummy
	var tmpPtr *main.ListNode
	p1, p2 := leftNodes, rightNodes
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			tmp.Next = p1
			tmpPtr = p1.Next
			p1.Next = nil
			p1 = tmpPtr
		} else {
			tmp.Next = p2
			tmpPtr = p2.Next
			p2.Next = nil
			p2 = tmpPtr
		}
		tmp = tmp.Next
	}

	if p1 == nil {
		tmp.Next = p2
	} else {
		tmp.Next = p1
	}

	ans := dummy.Next
	dummy.Next = nil
	return ans
}


func mergeKLists23(lists []*main.ListNode) *main.ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return lists[0]
	}

	return merge23(lists, 0, l - 1)
}
