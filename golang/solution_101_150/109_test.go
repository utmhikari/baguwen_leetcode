package solution_101_150

import "github.com/utmhikari/go-leetcode"

// 109. 有序链表转换二叉搜索树
// 给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
//
//本题中，一个高度平衡二叉树是指一个二叉树每个节点的左右两个子树的高度差的绝对值不超过 1。



func convertListToBST(nodes []*main.ListNode, left int, right int) *main.TreeNode {
	if right < left {
		return nil
	}

	mid := left + (right - left) / 2
	root := &main.TreeNode{
		Val: nodes[mid].Val,
	}

	root.Left = convertListToBST(nodes, left, mid - 1)
	root.Right = convertListToBST(nodes, mid + 1, right)

	return root
}


func sortedListToBST109(head *main.ListNode) *main.TreeNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &main.TreeNode{
			Val: head.Val,
		}
	}

	nodes := make([]*main.ListNode, 0)
	p := head
	for {
		nodes = append(nodes, p)
		p = p.Next
		if p == nil {
			break
		}
	}

	return convertListToBST(nodes, 0, len(nodes) - 1)
}
