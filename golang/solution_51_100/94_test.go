package solution_51_100

import "github.com/utmhikari/go-leetcode"

// 94. 二叉树的中序遍历


var ans94 []int


func traverse94(root *main.TreeNode) {
	if root == nil {
		return
	}
	traverse94(root.Left)
	ans94 = append(ans94, root.Val)
	traverse94(root.Right)
}

func inorderTraversal94(root *main.TreeNode) []int {
	ans94 = make([]int, 0)
	traverse94(root)
	return ans94
}
