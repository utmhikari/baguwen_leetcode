from node import TreeNode


def visit(r: TreeNode, a: [], d: int):
    if not r:
        return
    if len(a) <= d:
        a.append([])
    a[d].append(r.val)
    visit(r.left, a, d + 1)
    visit(r.right, a, d + 1)


class Solution:
    def levelOrder(self, root: TreeNode) -> [[int]]:
        a = []
        visit(root, a, 0)
        return a
