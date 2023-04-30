package solution_101_150

import "github.com/utmhikari/go-leetcode"

//105. 从前序与中序遍历序列构造二叉树


func solve105(preorder []int, p1 int, p2 int, inorder []int, i1 int, i2 int) *main.TreeNode {
	if p2 < p1 {
		return nil
	}
	if p2 == p1 {
		return &main.TreeNode{Val: preorder[p1]}
	}

	i := i1
	for i <= i2 {
		if inorder[i] == preorder[p1] {
			break
		}
		i++
	}

	root := &main.TreeNode{Val: preorder[p1]}
	root.Left = solve105(preorder, p1 + 1, p1 + (i - i1), inorder, i1, i - 1)
	root.Right = solve105(preorder, p1 + (i + 1 - i1), p2, inorder, i + 1, i2)
	return root
}



func buildTree(preorder []int, inorder []int) *main.TreeNode {
	lp, li := len(preorder), len(inorder)
	if lp != li || lp == 0 {
		return nil
	}
	if lp == 1 {
		return &main.TreeNode{Val: preorder[0]}
	}
	return solve105(preorder, 0, lp - 1, inorder, 0, li - 1)
}