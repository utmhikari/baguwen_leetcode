package solution_201_250

import "github.com/utmhikari/go-leetcode"

//237. 删除链表中的节点
//请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点。传入函数的唯一参数为 要被删除的节点 。


func deleteNode237(node *main.ListNode) {
	if node.Next != nil {
		node.Val = node.Next.Val
		node.Next = node.Next.Next
	} else {
		node = nil
	}
}