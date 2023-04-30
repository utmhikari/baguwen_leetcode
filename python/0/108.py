from node import TreeNode

def satb(ns, l, r):
    if l > r:
        return None
    if l == r:
        return TreeNode(ns[l])
    mid = l + (r - l) // 2
    root = TreeNode(ns[mid])
    root.left = satb(ns, l, mid - 1)
    root.right = satb(ns, mid + 1, r)
    return root

class Solution:
    def sortedArrayToBST(self, nums: [int]) -> TreeNode:
        return satb(nums, 0, len(nums) - 1)
