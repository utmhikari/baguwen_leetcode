class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1:
            return s
        store = []
        for _ in range(numRows):
            store.append([])
        i = -1
        p = True
        for c in s:
            if p:
                i += 1
            else:
                i -= 1
            store[i].append(c)
            if i == 0:
                p = True
            elif i == numRows - 1:
                p = False
        ans = ""
        for cs in store:
            for c in cs:
                ans += c
        return ans


so = Solution()

s = "LEETCODEISHIRING"

"""
l   c   i   r
e t o e s i i g
e   d   h   n
"""
print(so.convert(s, 3))


"""
l     d     r
e   o e   i i
e c   i h   n
t     s     g
"""
print(so.convert(s, 4))