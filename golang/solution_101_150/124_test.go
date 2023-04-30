package solution_101_150

import (
	"github.com/utmhikari/go-leetcode"
	"math"
)

//124. 二叉树中的最大路径和
//路径 被定义为一条从树中任意节点出发，
//沿父节点-子节点连接，达到任意节点的序列。
//同一个节点在一条路径序列中 至多出现一次 。
//该路径 至少包含一个 节点，且不一定经过根节点。
//
//路径和 是路径中各节点值的总和。
//
//给你一个二叉树的根节点 root ，返回其 最大路径和


// solve124 returns maxValWithRoot, maxValOnRoot
func solve124(root *main.TreeNode) (int, int) {
	if nil == root {
		return math.MinInt32, math.MinInt32  // cannot reach it
	}
	leftMaxValWithRoot, leftMaxValOnRoot := solve124(root.Left)
	rightMaxValWithRoot, rightMaxValOnRoot := solve124(root.Right)

	// cal maxValWithRoot
	maxValWithRoot := root.Val
	if leftMaxValWithRoot > 0 && rightMaxValWithRoot > 0 {
		if leftMaxValWithRoot > rightMaxValWithRoot {
			maxValWithRoot += leftMaxValWithRoot
		} else {
			maxValWithRoot += rightMaxValWithRoot
		}
	} else if leftMaxValWithRoot > 0 {
		maxValWithRoot += leftMaxValWithRoot
	} else if rightMaxValWithRoot > 0 {
		maxValWithRoot += rightMaxValWithRoot
	}

	// cal maxValOnRoot
	maxValOnRoot := root.Val
	if leftMaxValWithRoot > 0 {
		maxValOnRoot += leftMaxValWithRoot
	}
	if rightMaxValWithRoot > 0 {
		maxValOnRoot += rightMaxValWithRoot
	}
	if leftMaxValOnRoot > maxValOnRoot {
		maxValOnRoot = leftMaxValOnRoot
	}
	if rightMaxValOnRoot > maxValOnRoot {
		maxValOnRoot = rightMaxValOnRoot
	}
	return maxValWithRoot, maxValOnRoot
}


func maxPathSum124(root *main.TreeNode) int {
	if nil == root {
		return 0
	}
	maxValWithRoot, maxValOnRoot := solve124(root)
	if maxValOnRoot > maxValWithRoot {
		return maxValOnRoot
	} else {
		return maxValWithRoot
	}
}
