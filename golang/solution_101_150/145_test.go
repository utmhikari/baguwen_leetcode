package solution_101_150

import "github.com/utmhikari/go-leetcode"

// 145. 二叉树的后序遍历

var ans145 []int

func solve145(root *main.TreeNode) {
	if root == nil {
		return
	}

	solve145(root.Left)
	solve145(root.Right)
	ans145 = append(ans145, root.Val)
}

func postorderTraversal145(root *main.TreeNode) []int {
	ans145 = make([]int, 0)
	solve145(root)
	return ans145
}