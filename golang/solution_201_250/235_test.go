package solution_201_250

import "github.com/utmhikari/go-leetcode"

//235. 二叉搜索树的最近公共祖先
//给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
//
//百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，
//最近公共祖先表示为一个结点 x，
//满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
//
//例如，给定如下二叉搜索树: root =[6,2,8,0,4,7,9,null,null,3,5]


func solve235(root, p, q *main.TreeNode) *main.TreeNode {
	if p == root || q == root {
		return root
	}
	if (root.Val > p.Val && root.Val < q.Val) || (root.Val < p.Val && root.Val > q.Val) {
		return root
	}
	if p.Val < root.Val {
		return solve235(root.Left, p, q)
	} else {
		return solve235(root.Right, p, q)
	}
}


func lowestCommonAncestor235(root, p, q *main.TreeNode) *main.TreeNode {
	return solve235(root, p, q)
}
