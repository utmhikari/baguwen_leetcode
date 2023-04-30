package solution_101_150


//116. 填充每个节点的下一个右侧节点指针
//
//给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
//
//struct Node {
//int val;
//Node *left;
//Node *right;
//Node *next;
//}
//填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
//
//初始状态下，所有next 指针都被设置为 NULL。


//func connect116(root *Node) *Node {
//	if root != nil && root.Left != nil {
//		root.Left.Next = root.Right
//		connect116(root.Left)
//		connect116(root.Right)
//		p1, p2 := root.Left, root.Right
//		if p1 != nil {
//			for {
//				p1 = p1.Right
//				p2 = p2.Left
//				if p1 != nil && p2 != nil {
//					p1.Next = p2
//				} else {
//					break
//				}
//			}
//		}
//
//	}
//
//	return root
//}