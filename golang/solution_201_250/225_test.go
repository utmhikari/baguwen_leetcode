package solution_201_250


//225. 用队列实现栈
//请你仅使用两个队列实现一个后入先出（LIFO）的栈，
//并支持普通队列的全部四种操作（push、top、pop 和 empty）。


type Q215 struct {
	arr []int
}

func ConstructorQ215() *Q215 {
	return &Q215{
		arr: make([]int, 0),
	}
}

func (this *Q215) PushToBack(x int) {
	this.arr = append([]int{x}, this.arr...)
}

func (this *Q215) PopFromFront() int {
	i := this.Top()
	this.arr = this.arr[:len(this.arr) - 1]
	return i
}

func (this *Q215) Top() int {
	return this.arr[len(this.arr) - 1]
}

func (this *Q215) Empty() bool {
	return len(this.arr) == 0
}

type MyStack struct {
	q1 *Q215 // pushToBack, popFromFront, top, empty
	q2 *Q215 // pushToBack, popFromFront, top, empty
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		q1: ConstructorQ215(),
		q2: ConstructorQ215(),
	}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	for !this.q1.Empty() {
		this.q2.PushToBack(this.q1.PopFromFront())
	}
	this.q1.PushToBack(x)
	for !this.q2.Empty() {
		this.q1.PushToBack(this.q2.PopFromFront())
	}
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	return this.q1.PopFromFront()
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.q1.Top()
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.q1.Empty()
}