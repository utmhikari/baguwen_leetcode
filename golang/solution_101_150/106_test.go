package solution_101_150

import "github.com/utmhikari/go-leetcode"

//106. 从中序与后序遍历序列构造二叉树
//根据一棵树的中序遍历与后序遍历构造二叉树。

// 中序：左上右
// 后序：左右上

func buildTree106Internal(inorder[] int, postorder []int,
	leftInOrder int, rightInOrder int, leftPostOrder int, rightPostOrder int) *main.TreeNode {
	if rightInOrder < leftInOrder || rightPostOrder < leftPostOrder {
		return nil
	}

	node := &main.TreeNode{
		Val: postorder[rightPostOrder],
	}

	i := leftInOrder
	for ; i <= rightInOrder; i++ {
		if inorder[i] == node.Val {
			break
		}
	}

	//fmt.Printf("%d\n", node.Val)
	node.Left = buildTree106Internal(inorder, postorder, leftInOrder, i - 1, leftPostOrder, leftPostOrder + i - leftInOrder - 1)
	node.Right = buildTree106Internal(inorder, postorder, i + 1, rightInOrder, leftPostOrder + i - leftInOrder, rightPostOrder - 1)
	return node
}

func buildTree106(inorder []int, postorder []int) *main.TreeNode {
	if len(inorder) == 0 || len(inorder) != len(postorder) {
		return nil
	}

	return buildTree106Internal(inorder, postorder, 0, len(inorder) - 1, 0, len(postorder) - 1)
}
