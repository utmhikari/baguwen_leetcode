package solution_351_400

import (
	"math/rand"
)

//380. O(1) 时间插入、删除和获取随机元素
//设计一个支持在平均 时间复杂度 O(1) 下，执行以下操作的数据结构。
//
//insert(val)：当元素 val 不存在时，向集合中插入该项。
//remove(val)：元素 val 存在时，从集合中移除该项。
//getRandom：随机返回现有集合中的一项。每个元素应该有相同的概率被返回。


type RandomizedSet struct {
	mp map[int]int
	ls []int
}


/** Initialize your data structure here. */
func Constructor380() RandomizedSet {
	return RandomizedSet{
		mp: make(map[int]int),
		ls: make([]int, 0),
	}
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.mp[val]
	if ok {
		return false
	}

	this.mp[val] = len(this.ls)
	this.ls = append(this.ls, val)
	return true
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.mp[val]
	if !ok {
		return false
	}

	l := len(this.ls)
	this.ls[idx], this.ls[l - 1] = this.ls[l - 1], this.ls[idx]
	this.mp[this.ls[idx]] = idx
	delete(this.mp, val)
	this.ls = this.ls[:l - 1]
	return true
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	if len(this.ls) == 0 {
		return 0
	}
	return this.ls[rand.Intn(len(this.ls))]
}
