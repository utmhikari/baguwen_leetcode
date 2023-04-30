package solution_201_250


//232. 用栈实现队列
// push, pop, peek, empty

type Stack232 struct {
	arr []int
}

func CtorStack232() *Stack232 {
	return &Stack232{arr: make([]int, 0)}
}

func (this *Stack232) Push(i int) {
	this.arr = append(this.arr, i)
}

func (this *Stack232) Pop() int {
	n := this.Peek()
	this.arr = this.arr[:len(this.arr) - 1]
	return n
}

func (this *Stack232) Peek() int {
	return this.arr[len(this.arr) - 1]
}

func (this *Stack232) Empty() bool {
	return len(this.arr) == 0
}


type MyQueue struct {
	stk1 *Stack232
	stk2 *Stack232
}


/** Initialize your data structure here. */
func Constructor232() MyQueue {
	return MyQueue{
		stk1: CtorStack232(),
		stk2: CtorStack232(),
	}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	for !this.stk1.Empty() {
		this.stk2.Push(this.stk1.Pop())
	}
	this.stk1.Push(x)
	for !this.stk2.Empty() {
		this.stk1.Push(this.stk2.Pop())
	}
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	return this.stk1.Pop()
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.stk1.Peek()
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.stk1.Empty()
}