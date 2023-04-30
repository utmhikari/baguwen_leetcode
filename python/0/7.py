import math

class Solution:
    def reverse(self, x: int) -> int:
        if x == 0:
            return 0
        mx = math.floor(pow(2, 31))
        y = x
        minus = False
        if y < 0:
            minus = True
            y = -y
        ans = 0
        while y != 0:
            ans = ans * 10 + (y % 10)
            y //= 10
        if minus:
            if ans > mx:
                return 0
            else:
                return -ans
        else:
            if ans > mx - 1:
                return 0
            else:
                return ans


s = Solution()
a = 123
print(s.reverse(a))

a = -123
print(s.reverse(a))

a = 120
print(s.reverse(a))

a = 1234567889
print(s.reverse(a))