# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


from node import TreeNode



class Solution:
    def addOneRow(self, root: TreeNode, v: int, d: int) -> TreeNode:
        if d <= 1:
            t = TreeNode(v)
            if d == 1:
                t.left = root
            else:
                t.right = root
            return t
        if root and d > 1:
            if d == 2:
                root.left = self.addOneRow(root.left, v, 1)
                root.right = self.addOneRow(root.right, v, 0)
            else:
                root.left = self.addOneRow(root.left, v, d - 1)
                root.right = self.addOneRow(root.right, v, d - 1)
        return root
