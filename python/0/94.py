from node import TreeNode

class Solution:
    def inorderTraversal(self, root: TreeNode):
        if not root:
            return [None]
        p = root
        a = []
        s = []
        while p or len(s) > 0:
            while p:
                s.append(p)
                p = p.left
            p = s.pop()
            a.append(p.val)
            p = p.right
        return a





s = Solution()

