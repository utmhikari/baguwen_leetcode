package solution_51_100

import "github.com/utmhikari/go-leetcode"

//100. 相同的树

func solve100(p *main.TreeNode, q *main.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return solve100(p.Left, q.Left) && solve100(p.Right, q.Right)
}


func isSameTree100(p *main.TreeNode, q *main.TreeNode) bool {
	return solve100(p, q)
}