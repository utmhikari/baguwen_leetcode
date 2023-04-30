package solution_101_150

import "github.com/utmhikari/go-leetcode"

//101. 对称二叉树
//给定一个二叉树，检查它是否是镜像对称的。


func solve101(r1 *main.TreeNode, r2 *main.TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 == nil || r2 == nil {
		return false
	}
	if r1.Val != r2.Val {
		return false
	}
	return solve101(r1.Left, r2.Right) && solve101(r1.Right, r2.Left)
}

func isSymmetric101(root *main.TreeNode) bool {
	if nil == root {
		return true
	}
	return solve101(root.Left, root.Right)
}