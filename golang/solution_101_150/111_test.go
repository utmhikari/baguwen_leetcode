package solution_101_150

import "github.com/utmhikari/go-leetcode"

//111.给定一个二叉树，找出其最小深度。
//
//最小深度是从根节点到最近叶子节点的最短路径上的节点数量。


func minInt(i1 int, i2 int) int {
	if i1 < i2 {
		return i1
	}

	return i2
}


func solveMinDepth(root *main.TreeNode, prevDepth int) int {
	if root.Left == nil && root.Right == nil {
		return prevDepth
	}

	d1, d2 := -1, -1

	if root.Left != nil {
		d1 = solveMinDepth(root.Left, prevDepth + 1)
	}

	if root.Right != nil {
		d2 = solveMinDepth(root.Right, prevDepth + 1)
	}

	if d1 == -1 {
		return d2
	}

	if d2 == -1 {
		return d1
	}

	return minInt(d1, d2)
}


func minDepth(root *main.TreeNode) int {
	if root == nil {
		return 0
	}

	return solveMinDepth(root, 1)
}