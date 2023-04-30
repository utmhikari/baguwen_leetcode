class Solution:
    def firstUniqChar(self, s: str) -> int:
        l = len(s)
        if l == 0:
            return -1
        cs = set()
        dups = set()
        for c in s:
            if c not in cs:
                cs.add(c)
            else:
                dups.add(c)
        for i in range(l):
            if s[i] not in dups:
                return i
        return -1