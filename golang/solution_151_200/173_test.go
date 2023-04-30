package solution_151_200

import "github.com/utmhikari/go-leetcode"

//173. 二叉搜索树迭代器
//实现一个二叉搜索树迭代器你将使用二叉搜索树的根节点初始化迭代器
//
//调用 next() 将返回二叉搜索树中的下一个最小的数


type BSTIterator struct {
	vals []int
}


func Constructor173(root *main.TreeNode) BSTIterator {
	it := BSTIterator{
		vals: make([]int, 0),
	}

	it.push(root)
	return it
}


func (it *BSTIterator) Next() int {
	return it.pop()
}


func (it *BSTIterator) HasNext() bool {
	return len(it.vals) > 0
}


func (it *BSTIterator) push(node *main.TreeNode) {
	if node == nil {
		return
	}

	it.push(node.Left)
	it.vals = append(it.vals, node.Val)
	it.push(node.Right)
}

func (it *BSTIterator) pop() int {
	val := it.vals[0]
	it.vals = it.vals[1:]
	return val
}