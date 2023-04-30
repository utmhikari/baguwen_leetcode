package solution_351_400

import (
	"github.com/utmhikari/go-leetcode"
	"math/rand"
)

//382. 链表随机节点

//给定一个单链表，随机选择链表的一个节点，并返回相应的节点值。保证每个节点被选的概率一样。


type Solution382 struct {
	l int
	heads []*main.ListNode

	cap int
}


/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor382(head *main.ListNode) Solution382 {
	s := Solution382{
		l:     0,
		heads: make([]*main.ListNode, 0),
		cap: 10000,
	}

	p := head
	for p != nil {
		if s.l % s.cap == 0 {
			s.heads = append(s.heads, p)
		}
		s.l++
		p = p.Next
	}
	return s
}


/** Returns a random node's value. */
func (s *Solution382) GetRandom() int {
	idx := rand.Intn(s.l)
	div, mod := idx / s.cap, idx % s.cap
	p := s.heads[div]
	for i := 0; i < mod; i++ {
		p = p.Next
	}
	return p.Val
}