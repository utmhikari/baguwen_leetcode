package solution_151_200

import "github.com/utmhikari/go-leetcode"

//199. 二叉树的右视图
//给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。


var ans199 []int


func visit199(root *main.TreeNode, depth int) {
	if nil == root {
		return
	}
	if len(ans199) < depth {
		ans199 = append(ans199, root.Val)
	}
	visit199(root.Right, depth + 1)
	visit199(root.Left, depth + 1)
}


func rightSideView199(root *main.TreeNode) []int {
	ans199 = make([]int, 0)
	visit199(root, 1)
	return ans199
}