package solution_201_250

import "github.com/utmhikari/go-leetcode"

var cnt230 int
var k230 int
var ans230 *main.TreeNode


func solve230(root *main.TreeNode) {
	if root == nil {
		return
	}
	solve230(root.Left)
	if ans230 != nil {
		return
	}
	cnt230++
	if cnt230 == k230 {
		ans230 = root
		return
	}
	solve230(root.Right)
}

func kthSmallest(root *main.TreeNode, k int) int {
	cnt230 = 0
	k230 = k
	ans230 = nil
	solve230(root)
	return ans230.Val
}

