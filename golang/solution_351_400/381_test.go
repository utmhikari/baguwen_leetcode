package solution_351_400

import (
	"fmt"
	"math/rand"
)

//381. O(1) 时间插入、删除和获取随机元素 - 允许重复
//
//设计一个支持在平均 时间复杂度 O(1) 下， 执行以下操作的数据结构。
//
//注意: 允许出现重复元素。
//
//insert(val)：向集合中插入元素 val。
//remove(val)：当 val 存在时，从集合中移除一个 val。
//getRandom：从现有集合中随机获取一个元素。每个元素被返回的概率应该与其在集合中的数量呈线性相关。


type RandomizedCollection struct {
	mp map[int]map[int]bool
	ls []int
}


/** Initialize your data structure here. */
func Constructor381() RandomizedCollection {
	return RandomizedCollection{
		mp: make(map[int]map[int]bool),
		ls: make([]int, 0),
	}
}


/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (rc *RandomizedCollection) Insert(val int) bool {
	mpv, ok := rc.mp[val]
	l := len(rc.ls)
	if mpv == nil {
		rc.mp[val] = map[int]bool{l: true}
	} else {
		rc.mp[val][l] = true
	}
	rc.ls = append(rc.ls, val)
	return !ok
}


/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (rc *RandomizedCollection) Remove(val int) bool {
	_, ok := rc.mp[val]
	if !ok {
		return false
	}

	l := len(rc.ls)
	last := rc.ls[l - 1]
	if last == val {
		delete(rc.mp[val], l - 1)
	} else {
		for idx, _ := range rc.mp[val] {
			rc.ls[idx], rc.ls[l - 1] = rc.ls[l - 1], rc.ls[idx]
			delete(rc.mp[last], l - 1)
			rc.mp[last][idx] = true
			delete(rc.mp[val], idx)
			break
		}
	}

	rc.ls = rc.ls[:l - 1]
	if len(rc.mp[val]) == 0 {
		delete(rc.mp, val)
	}

	fmt.Printf("mp: %v, ls: %v\n", rc.mp, rc.ls)

	return true
}


/** Get a random element from the collection. */
func (rc *RandomizedCollection) GetRandom() int {
	l := len(rc.ls)
	if l == 0 {
		return 0
	}

	return rc.ls[rand.Intn(l)]
}