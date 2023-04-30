package solution_101_150

import "github.com/utmhikari/go-leetcode"

//114. 二叉树展开为链表
//
//给你二叉树的根结点 root ，请你将它展开为一个单链表：
//
//展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
//展开后的单链表应该与二叉树 先序遍历 顺序相同。


func getMostRight(root *main.TreeNode) *main.TreeNode {
	if root == nil {
		return root
	}

	tmp := root

	for {
		if tmp.Right == nil {
			return tmp
		}

		tmp = tmp.Right
	}
}


func flatten(root *main.TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	mostRight := getMostRight(root.Left)
	if mostRight != nil {
		mostRight.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
}
