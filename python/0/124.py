from node import TreeNode


# w: max for the parent
# wo: max in current root
def mps(root):
    if not root:
        return False, False
    w = root.val
    wo = root.val
    wl, wol = mps(root.left)
    wr, wor = mps(root.right)
    if not isinstance(wl, bool) or not isinstance(wr, bool):
        if isinstance(wl, bool):
            w = w + wr if wr > 0 else w
            wo = wo + wr if wo > 0 else wr
        elif isinstance(wr, bool):
            w = w + wl if wl > 0 else w
            wo = wo + wl if wo > 0 else wl
        else:
            w += max(0, wr, wl)
            wo = max(wr, wl, wr + root.val + wl, wr + root.val, wl + root.val)
    if not isinstance(wol, bool):
        wo = max(wo, wol)
    if not isinstance(wor, bool):
        wo = max(wo, wor)
    return w, wo

class Solution:
    def maxPathSum(self, root: TreeNode) -> int:
        w, wo = mps(root)
        return max(w, wo)

