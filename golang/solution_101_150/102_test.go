package solution_101_150

import "github.com/utmhikari/go-leetcode"

//102. 二叉树的层序遍历
//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。


var ans102 [][]int

func solve102(root *main.TreeNode, level int) {
	if root != nil {
		for len(ans102) < level {
			ans102 = append(ans102, make([]int, 0))
		}
		ans102[level - 1] = append(ans102[level - 1], root.Val)
		solve102(root.Left, level + 1)
		solve102(root.Right, level + 1)
	}

}


func levelOrder102(root *main.TreeNode) [][]int {
	ans102 = make([][]int, 0)
	if root != nil {
		solve102(root, 1)
	}
	return ans102
}