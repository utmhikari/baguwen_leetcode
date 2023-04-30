package solution_51_100

import (
	"github.com/utmhikari/go-leetcode"
	"math"
)

//98. 验证二叉搜索树


func isValid98(root *main.TreeNode, minValue int, maxValue int) bool {
	if root == nil || minValue > maxValue || root.Val < minValue || root.Val > maxValue {
		return false
	}

	if root.Left != nil {
		if !isValid98(root.Left, minValue, root.Val - 1) {
			return false
		}
	}

	if root.Right != nil {
		if !isValid98(root.Right, root.Val + 1, maxValue) {
			return false
		}
	}

	return true
}


func isValidBST98(root *main.TreeNode) bool {
	if nil == root {
		return true
	}
	return isValid98(root, math.MinInt32, math.MaxInt32)
}