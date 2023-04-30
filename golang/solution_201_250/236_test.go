package solution_201_250

import "github.com/utmhikari/go-leetcode"

//236. 二叉树的最近公共祖先


var ans236 *main.TreeNode


func solve236(root, p, q *main.TreeNode) (*main.TreeNode, *main.TreeNode) {
	if ans236 != nil {
		return nil, nil
	}
	if root == nil {
		return nil, nil
	}

	var p0, q0 *main.TreeNode
	if root == p {
		p0 = root
	} else if root == q {
		q0 = root
	}

	p1, q1 := solve236(root.Left, p, q)
	if p1 != nil && q1 != nil {
		if ans236 == nil {
			ans236 = root
		}
		return nil, nil
	} else if p1 != nil {
		p0 = p1
	} else if q1 != nil {
		q0 = q1
	}
	if p0 != nil && q0 != nil {
		ans236 = root
		return nil, nil
	}
	
	p1, q1 = solve236(root.Right, p, q)
	if p1 != nil && q1 != nil {
		if ans236 == nil {
			ans236 = root
		}
		return nil, nil
	} else if p1 != nil {
		p0 = p1
	} else if q1 != nil {
		q0 = q1
	}
	if p0 != nil && q0 != nil {
		ans236 = root
		return nil, nil
	}

	if p0 != nil {
		return p0, nil
	} else if q0 != nil {
		return nil, q0
	} else {
		return nil, nil
	}
}



func lowestCommonAncestor(root, p, q *main.TreeNode) *main.TreeNode {
	ans236 = nil
	solve236(root, p, q)
	return ans236
}