from node import TreeNode, printtree


def bt(p, i, lp, rp, li, ri):
    if lp > rp:
        return None
    if lp == rp:
        return TreeNode(p[lp])
    root = TreeNode(p[lp])
    k = li
    while k <= ri:
        if i[k] == p[lp]:
            break
        k += 1
    root.left = bt(p, i, lp + 1, lp + k - li, li, k - 1)
    root.right = bt(p, i, lp + k - li + 1, rp, k + 1, ri)
    return root


class Solution:
    def buildTree(self, preorder: [int], inorder: [int]) -> TreeNode:
        lp = len(preorder)
        li = len(inorder)
        return bt(preorder, inorder, 0, lp - 1, 0, li - 1)


s = Solution()
printtree(s.buildTree([3,9,20,15,7], [9,3,15,20,7]))

