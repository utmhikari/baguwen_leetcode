package solution_101_150

import "github.com/utmhikari/go-leetcode"

//129. 求根节点到叶节点数字之和
//给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
//每条从根节点到叶节点的路径都代表一个数字：
//
//例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
//计算从根节点到叶节点生成的 所有数字之和 。
//
//叶节点 是指没有子节点的节点。


var sum129 int

func traverse129(root *main.TreeNode, curSum int) {
	curSum = curSum * 10 + root.Val
	if root.Left == nil && root.Right == nil {
		sum129 += curSum
		return
	}

	if root.Left != nil {
		traverse129(root.Left, curSum)
	}

	if root.Right != nil {
		traverse129(root.Right, curSum)
	}
}


func sumNumbers129(root *main.TreeNode) int {
	if root == nil {
		return 0
	}
	sum129 = 0
	traverse129(root, 0)
	return sum129
}
