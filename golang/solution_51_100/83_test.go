package solution_51_100

import (
	"fmt"
	"github.com/utmhikari/go-leetcode"
	"testing"
)

// 82. 删除排序链表中的重复元素 II
// 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。


func deleteDuplicates83(head *main.ListNode) *main.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &main.ListNode{
		Val: -1,
		Next: head,
	}
	p0, p1, p2 := dummy, head, head.Next
	for {
		if p2 == nil {
			break
		}
		if p2.Val == p1.Val {
			for {
				if p2 == nil || p2.Val != p1.Val {
					break
				}
				p2 = p2.Next
			}

			p1.Next = p2
			p0.Next = p1
			p0 = p0.Next
			if p2 == nil {
				break
			}
			p1 = p2
			p2 = p2.Next
		} else {
			p0.Next = p1
			p1 = p2
			p2 = p2.Next
			p0 = p0.Next
		}
	}

	newHead := dummy.Next
	dummy.Next = nil
	return newHead
}


type test83 struct {
	nums	[]int
	expected []int
}


func Test_83 (t *testing.T) {
	cases := append(make([]test83, 0),
		test83{
			nums:     []int{1,1,2},
			expected: []int{1,2},
		},
		test83{
			nums:     []int{1,1,2,3,3},
			expected: []int{1,2,3},
		},
	)

	var ans *main.ListNode

	for _, c := range cases {
		ans = deleteDuplicates83(main.GenLinkedListNode(c.nums))
		if !ans.Equals(main.GenLinkedListNode(c.expected)) {
			t.Errorf("failed at %v+, ans: %s\n", c, ans.ToString())
		} else {
			fmt.Printf("success at %v+\n", c)
		}
	}
}
