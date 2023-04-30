package solution_51_100

import (
	"fmt"
	"github.com/utmhikari/go-leetcode"
)


//95. 不同的二叉搜索树 II
//给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。
//可以按 任意顺序 返回答案。



var map95 map[string][]*main.TreeNode


func solve95(left int, right int) []*main.TreeNode {
	if left > right {
		return []*main.TreeNode{}
	}
	if left == right {
		return []*main.TreeNode{{Val: left}}
	}
	key := fmt.Sprintf("%d_%d", left, right)
	ans, ok := map95[key]
	if ok {
		return ans
	}

	ans = make([]*main.TreeNode, 0)
	for i := left; i <= right; i++ {
		leftAnsList := solve95(left, i - 1)
		rightAnsList := solve95(i + 1, right)
		if len(leftAnsList) == 0 && len(rightAnsList) == 0 {
			ans = append(ans, &main.TreeNode{Val: i})
		} else if len(leftAnsList) == 0 {
			for _, rightAns := range rightAnsList {
				root := &main.TreeNode{Val: i}
				root.Right = rightAns
				ans = append(ans, root)
			}
		} else if len(rightAnsList) == 0 {
			for _, leftAns := range leftAnsList {
				root := &main.TreeNode{Val: i}
				root.Left = leftAns
				ans = append(ans, root)
			}
		} else {
			for _, leftAns := range leftAnsList {
				for _, rightAns := range rightAnsList {
					root := &main.TreeNode{Val: i}
					root.Left = leftAns
					root.Right = rightAns
					ans = append(ans, root)
				}
			}
		}
	}
	map95[key] = ans
	return ans
}

func generateTrees95(n int) []*main.TreeNode {
	if n == 0 {
		return make([]*main.TreeNode, 0)
	}
	if n == 1 {
		return append(make([]*main.TreeNode, 0), &main.TreeNode{
			Val: 1,
		})
	}
	map95 = make(map[string][]*main.TreeNode)
	return solve95(1, n)
}