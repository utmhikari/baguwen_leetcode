

pm = [
    5, 7, 11, 13, 17, 19, 23, 29, 31, 41,
    43, 47, 53, 59, 61, 67, 71, 73, 79, 83,
    89, 97, 101, 103, 107, 109
]
md = 2 ** 32

def h(s):
    ls = len(s)
    hs = ls
    l = len(pm)
    ord_ps = [pm[ord(c) - ord('a')] for c in s]
    fc = 1
    for o in ord_ps:
        fc *= o
        hs *= (o * o * o)
        hs %= md
    for i in range(ls):
        p = pm[i % l]
        hs = (hs * p + fc) % md
    return hs

class Solution:
    def groupAnagrams(self, strs: [str]) -> [[str]]:
        d = dict()
        for s in strs:
            hs = h(s)
            print(hs)
            if hs not in d.keys():
                d[hs] = []
            d[hs].append(s)
        return [d[k] for k in d.keys()]


s = Solution()
print(s.groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"]))