from node import TreeNode


def visit(r: TreeNode, a: [], d: int):
    if not r:
        return
    if len(a) <= d:
        a.append([])
    if d % 2 == 0:
        a[d].append(r.val)
    else:
        a[d].insert(0, r.val)
    visit(r.left, a, d + 1)
    visit(r.right, a, d + 1)


class Solution:
    def zigzagLevelOrder(self, root: TreeNode) -> [[int]]:
        a = []
        visit(root, a, 0)
        return a
