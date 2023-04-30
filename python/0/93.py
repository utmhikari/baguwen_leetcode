from typing import List


def dfs(ans, tmp, s, cnt, left, right):
    sm = 0
    if left > right:
        return
    if cnt == 3 and left <= right:
        is_zero = s[left] == '0'
        while left <= right:
            sm = sm * 10 + int(s[left])
            if sm >= 256 or (is_zero and left < right):
                return
            left += 1
        if sm < 256:
            tmp.append(str(sm))
            ans.append('.'.join(tmp))
            tmp.pop()
    else:
        is_zero = s[left] == '0'
        while left <= right:
            sm = sm * 10 + int(s[left])
            left += 1
            if sm < 256:
                tmp.append(str(sm))
                dfs(ans, tmp, s, cnt + 1, left, right)
                tmp.pop()
            else:
                break
            if is_zero:
                break


class Solution:
    def restoreIpAddresses(self, s: str) -> List[str]:
        ans = []
        tmp = []
        l = len(s)
        if l < 4 or l > 16:
            return ans
        dfs(ans, tmp, s, 0, 0, l - 1)
        return ans


s = Solution()
print(s.restoreIpAddresses("25525511135"))
print(s.restoreIpAddresses("010010"))
