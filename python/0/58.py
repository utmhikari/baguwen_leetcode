class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        r = s.rstrip()
        l = len(r)
        if l == 0:
            return 0
        for i in range(l):
            if s[l - 1 - i] == ' ':
                return i
        return l