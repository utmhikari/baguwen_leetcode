class Solution:
    def longestCommonPrefix(self, strs) -> str:
        l = len(strs)
        if l == 0:
            return ''
        p = strs[-1]
        for i in range(l - 1):
            si = strs[i]
            lp = len(p)
            li = len(si)
            j = 0
            while j < min(lp, li):
                if p[j] == si[j]:
                    j += 1
                else:
                    break
            if j == 0:
                return ''
            else:
                p = p[:j]
        return p
