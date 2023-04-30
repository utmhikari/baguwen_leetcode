class TNode:
    def __init__(self, val):
        self.val = val
        self.left = None
        self.right = None
        self.cnt = 0


def insert(r: TNode, n: TNode, ans: [int], i: int):
    if not r:
        r = n
    elif r.val >= n.val:
        r.cnt += 1
        r.left = insert(r.left, n, ans, i)
    else:
        ans[i] += r.cnt + 1
        r.right = insert(r.right, n, ans, i)
    return r


class Solution:
    def countSmaller(self, nums: [int]):
        l = len(nums)
        ans = [0] * l
        root = None
        i = l - 1
        while i >= 0:
            root = insert(root, TNode(nums[i]), ans, i)
            i -= 1
        return ans


s = Solution()
print(s.countSmaller([5,2,6,1]))