package solution_101_150

import "github.com/utmhikari/go-leetcode"

func hasCycle141(head *main.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	dummy := &main.ListNode{Next: head}
	slow, fast := dummy, dummy
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
		if fast == slow {
			return true
		}
	}
	return false
}