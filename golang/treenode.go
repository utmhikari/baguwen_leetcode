package main

import (
	"bytes"
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func (t *TreeNode) Equals(tt *TreeNode) bool {
	if t == nil && tt == nil {
		return true
	}
	if t == nil || tt == nil || t.Val != tt.Val {
		return false
	}

	return t.Left.Equals(tt.Left) && t.Right.Equals(tt.Right)
}

func treeNodeToString(root *TreeNode, buffer bytes.Buffer) {
	if root == nil {
		if buffer.Len() == 0 {
			buffer.WriteString("nil")
		} else {
			buffer.WriteString(",nil")
		}
		return
	}

	if buffer.Len() == 0 {
		buffer.WriteString(fmt.Sprintf("%d", root.Val))
	} else {
		buffer.WriteString(fmt.Sprintf(",%d", root.Val))
	}

	treeNodeToString(root.Left, buffer)
	treeNodeToString(root.Right, buffer)
}


func (t *TreeNode) ToString() string {
	var buffer bytes.Buffer
	treeNodeToString(t, buffer)
	return buffer.String()
}
