package solution_301_350

import (
	"strings"
)

//331. 验证二叉树的前序序列化
//序列化二叉树的一种方法是使用前序遍历。
//当我们遇到一个非空节点时，我们可以记录下这个节点的值。
//如果它是一个空节点，我们可以使用一个标记值记录，例如 #。


var nodeStrs331 []string
var curIdx331 int


func solve331() bool {
	l := len(nodeStrs331)

	// cannot get node
	if curIdx331 >= l {
		return false
	}

	// nullptr
	if nodeStrs331[curIdx331] == "#" {
		curIdx331++
		return true
	}

	// add self
	curIdx331++

	// left
	leftOk := solve331()
	if !leftOk {
		return false
	}

	// right
	rightOk := solve331()
	if !rightOk {
		return false
	}

	// ok
	return true
}


func isValidSerialization(preorder string) bool {
	nodeStrs331 = strings.Split(preorder, ",")
	curIdx331 = 0

	l := len(nodeStrs331)
	if l == 0 {
		return false
	}
	ok := solve331()
	if !ok {
		return false
	}

	return curIdx331 == l // any left?
}