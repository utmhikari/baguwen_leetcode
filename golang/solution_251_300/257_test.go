package solution_251_300

import (
	"bytes"
	"github.com/utmhikari/go-leetcode"
	"strconv"
)

//257. 二叉树的所有路径
//给定一个二叉树，返回所有从根节点到叶子节点的路径。


var ans257 []string
var buf257 bytes.Buffer


func solve257(root *main.TreeNode) {
	if root == nil {
		return
	}

	if buf257.Len() > 0 {
		buf257.WriteString("->")
	}
	rootValStr := strconv.Itoa(root.Val)
	buf257.WriteString(rootValStr)

	if root.Left == nil && root.Right == nil {
		ans257 = append(ans257, buf257.String())
	}

	solve257(root.Left)
	solve257(root.Right)

	buf257.Truncate(buf257.Len() - len(rootValStr))
	if buf257.Len() > 0 {
		buf257.Truncate(buf257.Len() - 2) // ->
	}
}


func binaryTreePaths257(root *main.TreeNode) []string {
	ans257 = make([]string, 0)
	buf257.Reset()
	solve257(root)
	return ans257
}