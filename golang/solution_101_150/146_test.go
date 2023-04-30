package solution_101_150

import (
	"bytes"
	"fmt"
)

type LRUNode146 struct {
	Key int
	Val int
	Next *LRUNode146
	Prev *LRUNode146
}


type LRUCache struct {
	MaxSize int

	cache map[int]*LRUNode146
	sz int
	head *LRUNode146
	tail *LRUNode146
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache: make(map[int]*LRUNode146),
		sz: 0,
		MaxSize: capacity,
		head: nil,
		tail: nil,
	}
}

func (this *LRUCache) moveToHead(node *LRUNode146) *LRUNode146 {
	switch {
	case node.Key == this.head.Key:
		break
	case node.Key == this.tail.Key:
		node.Prev.Next = nil
		this.tail = node.Prev
		node.Prev = nil
		node.Next = this.head
		this.head.Prev = node
		this.head = node
		break
	default:
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		node.Prev = nil
		node.Next = this.head
		this.head.Prev = node
		this.head = node
		break
	}
	return node
}

func (this *LRUCache) appendToHead(node *LRUNode146) {
	if this.sz == 0 {
		this.head = node
		this.tail = node
	} else {
		node.Next = this.head
		this.head.Prev = node
		this.head = node
	}
	this.sz++
}

func (this *LRUCache) removeFromTail() {
	if this.sz == 0 {
		return
	}
	tail := this.tail
	if this.sz == 1 {
		this.head = nil
		this.tail = nil
	} else {
		tail.Prev.Next = nil
		this.tail = tail.Prev
		tail.Prev = nil
	}
	delete(this.cache, tail.Key)
	this.sz--
}

func (this *LRUCache) Dump() {
	if this.head == nil {
		fmt.Println("Dump: nil")
	}
	p := this.head
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("(%d,%d)", p.Key, p.Val))
	for p.Next != nil {
		b.WriteString("->")
		p = p.Next
		b.WriteString(fmt.Sprintf("(%d,%d)", p.Key, p.Val))
	}
	fmt.Printf("Dump: %s\n", b.String())
}

func (this *LRUCache) DumpReverse() {
	if this.tail == nil {
		fmt.Println("DumpReverse: nil")
	}
	p := this.tail
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("(%d,%d)", p.Key, p.Val))
	for p.Prev != nil {
		b.WriteString("->")
		p = p.Prev
		b.WriteString(fmt.Sprintf("(%d,%d)", p.Key, p.Val))
	}
	fmt.Printf("DumpReverse: %s\n", b.String())
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	this.Dump()
	this.DumpReverse()
	return node.Val
}


func (this *LRUCache) Put(key int, value int)  {
	node, ok := this.cache[key]
	if ok {
		node.Val = value
		this.moveToHead(node)
	} else {
		if this.sz == this.MaxSize {
			this.removeFromTail()
		}
		node = &LRUNode146{Key: key, Val: value}
		this.cache[key] = node
		this.appendToHead(node)
	}
	this.Dump()
	this.DumpReverse()
}