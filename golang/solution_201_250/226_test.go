package solution_201_250

import "github.com/utmhikari/go-leetcode"

//226. 翻转二叉树


func invertTree(root *main.TreeNode) *main.TreeNode {
	if root != nil {
		root.Left = invertTree(root.Left)
		root.Right = invertTree(root.Right)

		root.Left, root.Right = root.Right, root.Left
	}

	return root
}