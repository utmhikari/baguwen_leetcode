# 1 <= steps <= 500
# 1 <= arrLen <= 10^6
import json


def solve(ans, steps, length, pos):
    if pos > steps or pos >= length or pos < 0:
        return 0
    elif pos == steps:
        return 1
    if pos not in ans.keys():
        ans[pos] = dict()
    if steps not in ans[pos].keys():
        ans[pos][steps] = solve(ans, steps - 1, length, pos - 1) + \
                          solve(ans, steps - 1, length, pos) + \
                          solve(ans, steps - 1, length, pos + 1)
    return ans[pos][steps] % (pow(10, 9) + 7)



class Solution:
    def numWays(self, steps: int, arrLen: int) -> int:
        ans = {}
        a = solve(ans, steps, arrLen, 0)
        # print(json.dumps(ans, indent=2))
        return a



s = Solution()
print(s.numWays(3, 2))
print(s.numWays(2, 4))
print(s.numWays(4, 2))
print(s.numWays(4, 3))
print(s.numWays(6, 13))
