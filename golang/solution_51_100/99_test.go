package solution_51_100

import "github.com/utmhikari/go-leetcode"

//99. 恢复二叉搜索树
//给你二叉搜索树的根节点 root ，该树中的两个节点被错误地交换。
//请在不改变其结构的情况下，恢复这棵树。
//进阶：使用 O(n) 空间复杂度的解法很容易实现。你能想出一个只使用常数空间的解决方案吗




func recoverTree99(root *main.TreeNode)  {
	//Morris 中序遍历
	var x, y, pred, predecessor *main.TreeNode

	// pred is the prev one against root

	for root != nil {
		if root.Left != nil {
			predecessor = root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				// first time find the rightmost of root.left
				predecessor.Right = root
				// root became next left
				root = root.Left
			} else {
				// the root has already been marked, compare
				if pred != nil && root.Val < pred.Val {
					y = root
					if x == nil {
						x = pred
					}
				}
				pred = root
				predecessor.Right = nil
				root = root.Right
			}
		} else {
			// no root.left anymore, try the bottom-most of root.right
			if pred != nil && root.Val < pred.Val {
				y = root
				if x == nil {
					x = pred
				}
			}
			pred = root
			// if originally root.Right is nil, it should point back to root itself
			root = root.Right
		}

	}
	x.Val, y.Val = y.Val, x.Val
}