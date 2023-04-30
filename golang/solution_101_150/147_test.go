package solution_101_150

import (
	"github.com/utmhikari/go-leetcode"
	"math"
)

//147. 对链表进行插入排序
//对链表进行插入排序。
//
//
//插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
//每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。
//
//插入排序算法：
//
//插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
//每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
//重复直到所有输入数据插入完为止。




func insertionSortList147(head *main.ListNode) *main.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &main.ListNode{
		Next: head,
	}
	p := dummy
	var minVal int
	var p0 *main.ListNode
	var p1 *main.ListNode
	var cur *main.ListNode
	var curPrev *main.ListNode

	for p != nil && p.Next != nil {
		minVal = math.MaxInt32
		p1 = p
		p0 = p.Next

		for {
			if p0.Val < minVal {
				minVal = p0.Val
				cur = p0
				curPrev = p1
			}

			if p0.Next == nil {
				break
			}

			p1 = p1.Next
			p0 = p0.Next
		}

		if curPrev != nil && cur != nil {
			curPrev.Next = cur.Next
			cur.Next = p.Next
			p.Next = cur
			p = p.Next
		}
	}

	newHead := dummy.Next
	dummy.Next = nil
	return newHead
}