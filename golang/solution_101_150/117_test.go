package solution_101_150


//
//117. 填充每个节点的下一个右侧节点指针 II
//给定一个二叉树
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

// use official ans

//因为必须处理树上的所有节点，所以无法降低时间复杂度，但是可以尝试降低空间复杂度。
//
//在方法一中，因为对树的结构一无所知，所以使用队列保证有序访问同一层的所有节点，并建立它们之间的连接。然而不难发现：一旦在某层的节点之间建立了 \text{next}next 指针，那这层节点实际上形成了一个链表。因此，如果先去建立某一层的 \text{next}next 指针，再去遍历这一层，就无需再使用队列了。
//
//基于该想法，提出降低空间复杂度的思路：如果第 ii 层节点之间已经建立 \text{next}next 指针，就可以通过 \text{next}next 指针访问该层的所有节点，同时对于每个第 ii 层的节点，我们又可以通过它的 \rm leftleft 和 \rm rightright 指针知道其第 i+1i+1 层的孩子节点是什么，所以遍历过程中就能够按顺序为第 i + 1i+1 层节点建立 \text{next}next 指针。
//
//具体来说：
//
//从根节点开始。因为第 00 层只有一个节点，不需要处理。可以在上一层为下一层建立 \text{next}next 指针。该方法最重要的一点是：位于第 xx 层时为第 x + 1x+1 层建立 \text{next}next 指针。一旦完成这些连接操作，移至第 x + 1x+1 层为第 x + 2x+2 层建立 \text{next}next 指针。
//当遍历到某层节点时，该层节点的 \text{next}next 指针已经建立。这样就不需要队列从而节省空间。每次只要知道下一层的最左边的节点，就可以从该节点开始，像遍历链表一样遍历该层的所有节点。

//func connect117(root *Node) *Node {
//	start := root
//
//	for start != nil {
//		var nextStart, last *Node
//		handle := func(cur *Node) {
//			if cur == nil {
//				return
//			}
//			if nextStart == nil {
//				// nextStart is the most left of current level
//				nextStart = cur
//			}
//			if last != nil {
//				// last is the most right of current level, and connects nodes in current level
//				last.Next = cur
//			}
//			last = cur
//		}
//
//		for p := start; p != nil; p = p.Next {
//			// p.Next points to next one in current level
//			handle(p.Left)
//			handle(p.Right)
//		}
//
//		start = nextStart
//	}
//
//	return root
//}