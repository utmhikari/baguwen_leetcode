import math


class Solution:
    def myAtoi(self, string: str) -> int:
        mx, mn = math.floor(pow(2, 31)) - 1, -math.floor(pow(2, 31))
        ans = 0
        minus = 0
        parse = False
        for c in string.strip():
            if c < '0' or c > '9':
                if parse or minus != 0 or (c != '-' and c != '+'):
                    break
                parse = True
                if c == '-':
                    minus = 1
                else:
                    minus = 2
            else:
                parse = True
                ans = ans * 10 + int(c)
        if minus == 1:
            ans = -ans
        if ans > mx:
            return mx
        if ans < mn:
            return mn
        return ans



s = Solution()

print(s.myAtoi('42'))
print(s.myAtoi("   -42"))
print(s.myAtoi("4193 with words"))
print(s.myAtoi("words and 987"))
print(s.myAtoi("-91283472332"))
print(s.myAtoi("  -0012a42"))