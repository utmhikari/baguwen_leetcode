package solution_101_150

import "github.com/utmhikari/go-leetcode"

//112. 路径总和
//
//
//给你二叉树的根节点root 和一个表示目标和的整数targetSum
//判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和targetSum 。
//
//叶子节点 是指没有子节点的节点。




func findPathSum112(root *main.TreeNode, curSum int, targetSum int) bool {
	curSum += root.Val

	if root.Left == nil && root.Right == nil {
		return curSum == targetSum
	}

	if root.Left != nil {
		if findPathSum112(root.Left, curSum, targetSum) {
			return true
		}
	}

	if root.Right != nil {
		if findPathSum112(root.Right, curSum, targetSum) {
			return true
		}
	}

	return false
}


func hasPathSum112(root *main.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return findPathSum112(root, 0, targetSum)
}