a = {
    0: 0,
    1: 0
}


class Solution:
    def trailingZeroes(self, n: int) -> int:
        if n in a.keys():
            return a[n]
        tmp = n // 5
        ans = self.trailingZeroes(tmp) + tmp
        a[n] = ans
        return ans

s = Solution()
print(s.trailingZeroes(3))
print(s.trailingZeroes(5))

