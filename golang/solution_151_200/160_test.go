package solution_151_200

import "github.com/utmhikari/go-leetcode"

//160. 相交链表
//编写一个程序，找到两个单链表相交的起始节点。



func getIntersectionNode160(headA, headB *main.ListNode) *main.ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	la, lb := 0, 0
	pa, pb := headA, headB
	for pa != nil {
		la++
		pa = pa.Next
	}
	for pb != nil {
		lb++
		pb = pb.Next
	}
	pa, pb = headA, headB
	if la > lb {
		for i := 0; i < la - lb; i++ {
			pa = pa.Next
		}
	} else if lb > la {
		for i := 0; i < lb - la; i++ {
			pb = pb.Next
		}
	}
	for pa != nil && pb != nil {
		if pa == pb {
			return pa
		}
		pa = pa.Next
		pb = pb.Next
	}
	return nil
}