# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

from node import TreeNode



def visit(rt: TreeNode, lst: list, k: int):
    if len(lst) == k:
        return
    if rt.left:
        visit(rt.left, lst, k)
    if len(lst) != k:
        lst.append(rt.val)
    if len(lst) != k and rt.right:
        visit(rt.right, lst, k)


class Solution:
    def kthSmallest(self, root: TreeNode, k: int) -> int:
        lst = []
        visit(root, lst, k)
        return lst[-1]

