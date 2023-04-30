# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

from node import TreeNode
import sys



def isvb(root, mx, mn) -> bool:
    if not root:
        return True
    if root.val >= mx or root.val <= mn:
        return False
    if root.left:
        if not isvb(root.left, root.val, mn):
            return False
    if root.right:
        if not isvb(root.right, mx, root.val):
            return False
    return True


class Solution:
    def isValidBST(self, root: TreeNode) -> bool:
        return isvb(root, sys.maxsize, -sys.maxsize-1)


