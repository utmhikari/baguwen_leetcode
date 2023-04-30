package solution_101_150

import "github.com/utmhikari/go-leetcode"

//113. 路径总和 II
//
//给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
//
//叶子节点 是指没有子节点的节点。


var ans113 [][]int
var tmp113 []int

func appendAns() {
	tmpLen := len(tmp113)

	if tmpLen == 0 {
		return
	}

	newAns := make([]int, tmpLen)

	for i := 0; i < tmpLen; i++ {
		newAns[i] = tmp113[i]
	}

	ans113 = append(ans113, newAns)
}


func findPathSum113(root *main.TreeNode, curSum int, targetSum int) {
	curSum += root.Val
	tmp113 = append(tmp113, root.Val)

	if root.Left == nil && root.Right == nil {
		if curSum == targetSum {
			appendAns()
		}
		return
	}

	if root.Left != nil {
		findPathSum113(root.Left, curSum, targetSum)
		tmp113 = tmp113[:len(tmp113)-1]
	}

	if root.Right != nil {
		findPathSum113(root.Right, curSum, targetSum)
		tmp113 = tmp113[:len(tmp113)-1]
	}
}


func pathSum113(root *main.TreeNode, targetSum int) [][]int {
	ans113 = make([][]int, 0)
	if root == nil {
		return ans113
	}
	tmp113 = make([]int, 0)
	findPathSum113(root, 0, targetSum)
	return ans113
}
