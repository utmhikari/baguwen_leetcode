"""
输入：3
输出：
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
解释：
以上的输出对应以下 5 种不同结构的二叉搜索树：

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
"""

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from node import TreeNode
from typing import List
import copy


def genTree(left: int, right: int) -> List[TreeNode]:
    if left > right:
        return []
    if left == right:
        return [TreeNode(left)]
    trees: List[TreeNode] = []
    for i in range(left, right + 1):
        left_trees = genTree(left, i - 1)
        right_trees = genTree(i + 1, right)
        mid = TreeNode(i)
        if len(left_trees) == 0:
            for tree in right_trees:
                mid.right = tree
                trees.append(copy.deepcopy(mid))
        elif len(right_trees) == 0:
            for tree in left_trees:
                mid.left = tree
                trees.append(copy.deepcopy(mid))
        else:
            for lt in left_trees:
                mid.left = lt
                for rt in right_trees:
                    mid.right = rt
                    trees.append(copy.deepcopy(mid))
    return trees

class Solution:
    def generateTrees(self, n: int) -> List[TreeNode]:
        return genTree(1, n)
