from typing import List

d = {
    '': [],
    '2': ['a', 'b', 'c'],
    '3': ['d', 'e', 'f'],
    '4': ['g', 'h', 'i'],
    '5': ['j', 'k', 'l'],
    '6': ['m', 'n', 'o'],
    '7': ['p', 'q', 'r', 's'],
    '8': ['t', 'u', 'v'],
    '9': ['w', 'x', 'y', 'z']
}


def cb(s, l, r):
    if s in d.keys():
        return d[s]
    mid = (l + r) // 2
    llst = cb(s[:mid + 1], 0, mid)
    rlst = cb(s[mid + 1:], 0, r - mid - 1)
    lst = []
    for lstr in llst:
        for rstr in rlst:
            lst.append(lstr + rstr)
    d[s] = lst
    return lst


class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        l = len(digits)
        if l == 0:
            return []
        return cb(digits, 0, l - 1)


s = Solution()
print(s.letterCombinations('23'))
print(s.letterCombinations('5678'))