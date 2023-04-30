package solution_101_150

import "github.com/utmhikari/go-leetcode"

//110. 平衡二叉树
//给定一个二叉树，判断它是否是高度平衡的二叉树。
//本题中，一棵高度平衡二叉树定义为：一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。


func abs110(i1 int, i2 int) int {
	v := i1 - i2
	if v < 0 {
		v = -v
	}
	return v
}

func max110(i1 int, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

func solve110(root *main.TreeNode) (bool, int) {
	leftHeight, rightHeight, rootHeight := 0, 0, 0
	if root != nil {
		rootHeight++
		var ok bool
		if root.Left != nil {
			ok, leftHeight = solve110(root.Left)
			if !ok {
				return false, 0
			}
		}
		if root.Right != nil {
			ok, rightHeight = solve110(root.Right)
			if !ok {
				return false, 0
			}
		}
	}
	return abs110(leftHeight, rightHeight) <= 1, max110(leftHeight, rightHeight) + rootHeight
}


func isBalanced110(root *main.TreeNode) bool {
	if nil == root {
		return true
	}
	ok, _ := solve110(root)
	return ok
}

