
def solve(s1, s2) -> bool:
    l = len(s1)
    if l == 1:
        return s1 == s2
    d1, d2 = {}, {}
    for c in s1:
        if c not in d1.keys():
            d1[c] = 0
        d1[c] += 1
    for c in s2:
        if c not in d2.keys():
            d2[c] = 0
        d2[c] += 1
    if d1.keys() != d2.keys():
        return False
    for c in d1.keys():
        if d1[c] != d2[c]:
            return False
    for i in range(l - 1):
        s1l = s1[:i + 1]
        s1r = s1[i + 1:]
        s2l1 = s2[:i + 1]
        s2r1 = s2[i + 1:]
        if solve(s1l, s2l1) and solve(s1r, s2r1):
            return True
        s2l2 = s2[l - i - 1:]
        s2r2 = s2[:l - i - 1]
        if solve(s1l, s2l2) and solve(s1r, s2r2):
            return True
    return False


class Solution:
    def isScramble(self, s1: str, s2: str) -> bool:
        l1 = len(s1)
        l2 = len(s2)
        if l1 != l2:
            return False
        return solve(s1, s2)



s = Solution()
print(s.isScramble("great", "rgeat"))
print(s.isScramble("abcde", "caebd"))
print(s.isScramble("rgtae", "rgeat"))
print(s.isScramble("rgtae", "great"))
print(s.isScramble("etagr", "rgeat"))
print(s.isScramble("abcdefghijklmnopq", "efghijklmnopqcadb"))