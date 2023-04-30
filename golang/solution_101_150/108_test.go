package solution_101_150

import "github.com/utmhikari/go-leetcode"

//108. 将有序数组转换为二叉搜索树
//给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。
//高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。


func solve108(nums []int, left int, right int) *main.TreeNode {
	if right < left {
		return nil
	}
	if right == left {
		return &main.TreeNode{Val: nums[left]}
	}

	mid := left + (right - left) / 2
	root := &main.TreeNode{Val: nums[mid]}

	root.Left = solve108(nums, left, mid - 1)
	root.Right = solve108(nums, mid + 1, right)
	return root
}


func sortedArrayToBST108(nums []int) *main.TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}
	return solve108(nums, 0, l - 1)
}
