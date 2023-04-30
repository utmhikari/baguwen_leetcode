package solution_51_100

import (
	"fmt"
	"github.com/utmhikari/go-leetcode"
	"testing"
)

// 82. 删除排序链表中的重复元素 II
// 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。


func deleteDuplicates82(head *main.ListNode) *main.ListNode {
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

			p0.Next = p2
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


type test82 struct {
	nums	[]int
	expected []int
}


func Test_82 (t *testing.T) {
	cases := append(make([]test82, 0),
		test82{
			nums:     []int{1,2,3,3,4,4,5},
			expected: []int{1,2,5},
		},
		test82{
			nums:     []int{1,1,1,2,3},
			expected: []int{2,3},
		},
	)

	var ans *main.ListNode

	for _, c := range cases {
		ans = deleteDuplicates82(main.GenLinkedListNode(c.nums))
		if !ans.Equals(main.GenLinkedListNode(c.expected)) {
			t.Errorf("failed at %v+, ans: %s\n", c, ans.ToString())
		} else {
			fmt.Printf("success at %v+\n", c)
		}
	}
}
