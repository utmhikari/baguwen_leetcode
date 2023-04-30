package solution_1_50

import "github.com/utmhikari/go-leetcode"

//19.删除链表的倒数第 N 个结点
//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//进阶：你能尝试使用一趟扫描实现吗？


func removeNthFromEnd19(head *main.ListNode, n int) *main.ListNode {
	nodes := make([]*main.ListNode, 0)
	p := head
	for p != nil {
		nodes = append(nodes, p)
		p = p.Next
	}
	l := len(nodes)
	if n > l || l == 1 {
		return nil
	}

	switch n {
	case l:
		p = head.Next
		head.Next = nil
		return p
	case 1:
		nodes[l - 2].Next = nil
		break
	default:
		prev, next := nodes[l - 1 - n], nodes[l + 1 - n]
		prev.Next.Next = nil
		prev.Next = next
		break
	}
	return head
}
