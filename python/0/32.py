from collections import deque


class Solution:
    def longestValidParentheses(self, s: str) -> int:
        dq = deque()
        l = len(s)
        left, right = 0, l - 1
        while left < right and s[left] == ')':
            left += 1
        while right > left and s[right] == '(':
            right -= 1
        if right - left < 1:
            return 0
        mx = 0
        dct = dict()
        for i in range(left, right + 1):
            c = s[i]
            if c == '(':
                dq.append(i)
            elif len(dq) > 0:
                leftidx = dq.pop()
                # if can chain
                tmplen = i - leftidx + 1
                while leftidx > 0 and leftidx - 1 in dct.keys():
                    tmplen += (leftidx - 1) - dct[leftidx - 1] + 1
                    leftidx = dct[leftidx - 1]
                dct[i] = leftidx
                mx = max(tmplen, mx)
        return mx

s = Solution()
print(s.longestValidParentheses("(()"))
print(s.longestValidParentheses(")()())"))