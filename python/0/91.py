

m = dict()
d = dict()
for i in range(26):
    d[str(i + 1)] = chr(i + ord('A'))


def dec(s):
    if s in m.keys():
        return m[s]
    if s.startswith('0'):
        return 0
    if len(s) == 1:
        m[s] = 1
        return 1
    a = dec(s[1:])
    if len(s) == 2:
        a += 1 if s in d else 0
    else:
        a += dec(s[2:]) if s[:2] in d else 0
    m[s] = a
    return a


class Solution:
    def numDecodings(self, s: str) -> int:
        if len(s) == 0:
            return 0
        dec(s)
        return m[s] if s in m.keys() else 0


s = Solution()
print(s.numDecodings('12'))
print(s.numDecodings('226'))
