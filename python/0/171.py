class Solution:
    def titleToNumber(self, s: str) -> int:
        a = 0
        m = ord('A') - 1
        for c in s:
            a = a * 26 + ord(c) - m
        return a

s = Solution()
print(s.titleToNumber('A'))
print(s.titleToNumber('AB'))
print(s.titleToNumber('ZY'))