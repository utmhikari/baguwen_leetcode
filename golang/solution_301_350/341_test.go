package solution_301_350


//341. 扁平化嵌套列表迭代器
//给你一个嵌套的整型列表。请你设计一个迭代器，使其能够遍历这个整型列表中的所有整数。
//
//列表中的每一项或者为一个整数，或者是另一个列表。其中列表的元素也可能是整数或是其他列表。



type NestedInteger struct {
	nestedList []*NestedInteger
	isInt bool
	intValue int
}

func (this *NestedInteger) IsInteger() bool {
	return this.isInt
}

func (this *NestedInteger) GetInteger() int {
	return this.intValue
}


func (this *NestedInteger) SetInteger(v int) {
	this.intValue = v
	this.isInt = true
}

func (this *NestedInteger) Add(elem NestedInteger) {
	this.nestedList = append(this.nestedList, &elem)
}

func (this *NestedInteger) GetList() []*NestedInteger {
	return this.nestedList
}


type NestedIterator struct {
	stk []int
	cur int
}

func Constructor341(nestedList []*NestedInteger) *NestedIterator {
	it := NestedIterator{
		stk: make([]int, 0),
		cur: 0,
	}
	for _, nestedInt := range nestedList {
		it.add(nestedInt)
	}
	return &it
}

func (this *NestedIterator) add(nestedInt *NestedInteger) {
	if nestedInt.IsInteger() {
		this.stk = append(this.stk, nestedInt.GetInteger())
		return
	}

	for _, nextNestedInt := range nestedInt.GetList() {
		this.add(nextNestedInt)
	}
}

func (this *NestedIterator) Next() int {
	cur := this.cur
	this.cur++
	return this.stk[cur]

}

func (this *NestedIterator) HasNext() bool {
	return this.cur < len(this.stk)
}
