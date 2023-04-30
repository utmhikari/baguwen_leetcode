package solution_101_150

import "github.com/utmhikari/go-leetcode"

//142. 环形链表 II
//给定一个链表，返回链表开始入环的第一个节点。如果链表无环，则返回null。
//
//为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。


func detectCycle142(head *main.ListNode) *main.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	if head.Next == head {
		return head
	}

	dummy := &main.ListNode{
		Next: head,
	}
	p1, p2 := dummy, dummy
	cnt := 0

	for {
		p1 = p1.Next
		if p1.Next == nil {
			return nil
		}

		p2 = p2.Next
		if p2 == nil {
			return nil
		}
		p2 = p2.Next
		if p2 == nil {
			return nil
		}

		cnt++
		if p1 == p2 {
			break
		}
	}

	//fmt.Printf("%d\n", p1.Val)

	p3 := dummy.Next
	for {
		if p1.Next != p3 {
			p1 = p1.Next
			p3 = p3.Next
		} else {
			return p3
		}
	}
}