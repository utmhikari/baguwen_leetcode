package solution_201_250

import "github.com/utmhikari/go-leetcode"

//222. 完全二叉树的节点个数
//
//给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。
//
//完全二叉树 的定义如下：
//在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，
//并且最下面一层的节点都集中在该层最左边的若干位置。
//若最底层为第 h 层，则该层包含 1~ 2h 个节点。


func hasNode222(root *main.TreeNode, depth int, v int) bool {
	bits := 1 << (depth - 1)
	node := root
	for node != nil && bits > 0 {
		if bits & v == 0 {
			node = node.Left
		} else {
			node = node.Right
		}
		bits >>= 1
	}
	return node != nil
}

func countNodes222(root *main.TreeNode) int {
	if root == nil {
		return 0
	}

	node := root
	depth := 0
	for node.Left != nil {
		depth++
		node = node.Left
	}

	left, right := 1 << depth, (1 << (depth + 1)) - 1
	for left < right {
		mid := left + (right - left + 1) / 2
		if hasNode222(root, depth, mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return left
}