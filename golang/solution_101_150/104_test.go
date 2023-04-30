package solution_101_150

import "github.com/utmhikari/go-leetcode"

//104. 二叉树的最大深度

var ans104 int

func solve104(root *main.TreeNode, height int) {
	if root != nil {
		if root.Left == nil && root.Right == nil {
			if height > ans104 {
				ans104 = height
			}
		}
		solve104(root.Left, height + 1)
		solve104(root.Right, height + 1)
	}
}

func maxDepth104(root *main.TreeNode) int {
	ans104 = 0
	solve104(root, 1)
	return ans104
}