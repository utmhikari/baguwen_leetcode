

tb = []


def ad(adds: [str]) -> str:
    mxlen = 0
    for s in adds:
        mxlen = max(mxlen, len(s))
    ads = ''
    up = 0
    for i in range(mxlen):
        idx = -(i + 1)
        sm = 0
        for s in adds:
            if -idx <= len(s) and s[idx] != '0':
                sm += ord(s[idx]) - ord('0')
        sm += up
        up = sm // 10
        ads = str(sm % 10) + ads
    if up != 0:
        ads = str(up) + ads
    return ads


def mp(s1: str, s2: str, zerocnt: int) -> str:
    l1, l2 = len(s1), len(s2)
    if l1 == 1 and l2 == 1:
        return tb[ord(s1) - ord('0')][ord(s2) - ord('0')] + '0' * zerocnt
    adds = []
    if l1 != 1:
        for i in range(l1):
            adds.append(mp(s1[l1 - 1 - i], s2, i))
    else:
        for i in range(l2):
            adds.append(mp(s1, s2[l2 - 1 - i], i))
    return ad(adds) + '0' * zerocnt



class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        num1 = num1.lstrip('0')
        num2 = num2.lstrip('0')
        if len(num1) == 0 or len(num2) == 0:
            return '0'
        # collect table
        if len(tb) == 0:
            for i in range(10):
                tb.append([0] * 10)
            for i in range(10):
                for j in range(10):
                    tb[i][j] = str(i * j)
        # count zero
        p1, p2 = len(num1) - 1, len(num2) - 1
        zerocnt = 0
        while num1[p1] == '0':
            zerocnt += 1
            p1 -= 1
        while num2[p2] == '0':
            zerocnt += 1
            p2 -= 1
        return mp(num1[:p1 + 1], num2[:p2 + 1], 0) + '0' * zerocnt




s = Solution()
print(s.multiply('2', '3'))
print(s.multiply('123', '456'))
print(s.multiply('123456789', '987654321'))