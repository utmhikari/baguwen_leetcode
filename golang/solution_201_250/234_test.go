package solution_201_250

import "github.com/utmhikari/go-leetcode"

//234. 回文链表
//请判断一个链表是否为回文链表。


func reverse234(head *main.ListNode) *main.ListNode {
	var tail *main.ListNode
	p := head
	for p != nil {
		tmp := p.Next
		p.Next = tail
		tail = p
		p = tmp
	}
	return tail
}


func isPalindrome234(head *main.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	mid, last := head, head
	for {
		last = last.Next
		if last == nil {
			break
		}
		last = last.Next
		if last == nil {
			break
		}
		mid = mid.Next
	}

	right := mid.Next
	mid.Next = nil
	right = reverse234(right)

	ok := true
	p1, p2 := head, right
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			ok = false
			break
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	right = reverse234(right)
	mid.Next = right
	return ok
}