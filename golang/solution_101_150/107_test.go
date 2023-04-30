package solution_101_150

import "github.com/utmhikari/go-leetcode"

var ans107 [][]int


func traverse107(root *main.TreeNode, depth int) {
	if nil == root {
		return
	}

	if depth > len(ans107) {
		needLen := depth - len(ans107)
		newAns := make([][]int, depth)
		for i := 0; i < needLen; i++ {
			newAns[i] = make([]int, 0)
		}
		for i := needLen; i < depth; i++ {
			newAns[i] = ans107[i - needLen]
		}
		ans107 = newAns
	}

	//fmt.Printf("%v+, %d, %d\n", ans107, depth, root.Val)
	ans107[len(ans107) - depth] = append(ans107[len(ans107) - depth], root.Val)
	traverse107(root.Left, depth + 1)
	traverse107(root.Right, depth + 1)
}


func levelOrderBottom107(root *main.TreeNode) [][]int {
	ans107 = make([][]int, 0)
	traverse107(root, 1)
	return ans107
}

func levelOrderBottom107Official(root *main.TreeNode) [][]int {
	var levelOrder [][]int
	if root == nil {
		return levelOrder
	}
	var queue []*main.TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		var level []int
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		levelOrder = append(levelOrder, level)
	}
	for i := 0; i < len(levelOrder) / 2; i++ {
		levelOrder[i], levelOrder[len(levelOrder) - 1 - i] = levelOrder[len(levelOrder) - 1 - i], levelOrder[i]
	}
	return levelOrder
}