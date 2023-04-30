package solution_101_150

import "github.com/utmhikari/go-leetcode"

// 144. 二叉树的前序遍历

var ans144 []int


func solve144(root *main.TreeNode) {
	if root == nil {
		return
	}

	ans144 = append(ans144, root.Val)

	solve144(root.Left)
	solve144(root.Right)
}

func preorderTraversal144(root *main.TreeNode) []int {
	ans144 = make([]int, 0)
	solve144(root)
	return ans144
}