# 给定一个整数 n，求以1~n节点组成的二叉搜索树有多少种？
#
# 示例:
#
# 输入: 3
# 输出: 5
# 解释:
# 给定 n = 3, 一共有 5 种不同结构的二叉搜索树:
#
# 1         3     3      2      1
#  \       /     /      / \      \
#   3     2     1      1   3      2
#  /     /       \                 \
# 2     1         2                 3
from node import TreeNode


ANS = {
    0: 1,
    1: 1,
    2: 2,
    3: 5,
    4: 14,
    5: 42,
    6: 132,
    7: 429,
    8: 1430,
    9: 4862,
    10: 16796,
    11: 58786,
    12: 208012,
    13: 742900,
    14: 2674440,
    15: 9694845
}


class Solution:
    def numTrees(self, n: int) -> int:
        if n in ANS.keys():
            return ANS[n]
        num = 0
        for i in range(1, n + 1):
            num += self.numTrees(i - 1) * self.numTrees(n - i)
        return num


s = Solution()
for i in range(15, 1, -1):
    print(s.numTrees(i))
