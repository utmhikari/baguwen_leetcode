package solution_301_350

import "github.com/utmhikari/go-leetcode"

//337. 打家劫舍 III
//
//在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。
//这个地区只有一个入口，我们称之为“根”。
//除了“根”之外，每栋房子有且只有一个“父“房子与之相连。
//一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
//如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。
//
//计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。


// with root, without root
func solve337(root *main.TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	withLeft, withoutLeft := solve337(root.Left)
	withRight, withoutRight := solve337(root.Right)

	// with root
	withRoot := root.Val + withoutLeft + withoutRight

	// without root
	withoutRootLeft := withLeft
	if withoutRootLeft < withoutLeft {
		withoutRootLeft = withoutLeft
	}
	withoutRootRight := withRight
	if withoutRootRight < withoutRight {
		withoutRootRight = withoutRight
	}
	withoutRoot := withoutRootLeft + withoutRootRight

	return withRoot, withoutRoot
}


func rob337(root *main.TreeNode) int {
	withRoot, withoutRoot := solve337(root)
	if withRoot > withoutRoot {
		return withRoot
	}
	return withoutRoot
}