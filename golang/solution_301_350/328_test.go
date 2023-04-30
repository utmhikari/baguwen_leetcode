package solution_301_350

import "github.com/utmhikari/go-leetcode"

//328. 奇偶链表
//
//给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。
//请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。
//
//请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。



func solve328Sep(head *main.ListNode) *main.ListNode {
	if head == nil {
		return nil
	}
	odds := &main.ListNode{}
	evens := &main.ListNode{}

	// separate
	p1, p2 := odds, evens
	p, f := head, 1
	var tmp *main.ListNode
	for p != nil {
		tmp = p.Next
		p.Next = nil
		if f == 1 {
			p1.Next = p
			p1 = p1.Next
			f = 0
		} else {
			p2.Next = p
			p2 = p2.Next
			f = 1
		}
		p = tmp
	}

	newFirst := &main.ListNode{}
	pn := newFirst
	p = odds.Next
	odds.Next = nil
	for p != nil {
		tmp = p.Next
		p.Next = nil
		pn.Next = p
		pn = pn.Next
		p = tmp
	}
	p = evens.Next
	evens.Next = nil
	for p != nil {
		tmp = p.Next
		p.Next = nil
		pn.Next = p
		pn = pn.Next
		p = tmp
	}

	newHead := newFirst.Next
	newFirst.Next = nil
	return newHead
}



func oddEvenList(head *main.ListNode) *main.ListNode {
	return solve328Sep(head)
}
