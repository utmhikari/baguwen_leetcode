package solution_101_150

import "github.com/utmhikari/go-leetcode"

//103. 二叉树的锯齿形层序遍历
//给定一个二叉树，返回其节点值的锯齿形层序遍历。
//（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

var ans103 [][]int

func solve103(root *main.TreeNode, level int) {
	if root != nil {
		for len(ans103) < level {
			ans103 = append(ans103, make([]int, 0))
		}
		if level % 2 == 0 {
			ans103[level - 1] = append([]int{root.Val}, ans103[level - 1]...)
		} else {
			ans103[level - 1] = append(ans103[level - 1], root.Val)
		}
		solve103(root.Left, level + 1)
		solve103(root.Right, level + 1)
	}

}

func zigzagLevelOrder(root *main.TreeNode) [][]int {
	ans103 = make([][]int, 0)
	if root != nil {
		solve103(root, 1)
	}
	return ans103
}