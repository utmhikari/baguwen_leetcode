d = {
    1: '1',
    2: '11',
    3: '21',
    4: '1211',
    5: '111221',
}


class Solution:
    def countAndSay(self, n: int) -> str:
        if n in d.keys():
            return d[n]
        p = self.countAndSay(n - 1)
        cnt = 1
        tc = p[0]
        i = 1
        l = len(p)
        ns = ''
        while i < l:
            if p[i] == tc:
                cnt += 1
            else:
                ns += '%d%s' % (cnt, tc)
                cnt = 1
                tc = p[i]
            i += 1
        return ns + '%d%s' % (cnt, tc)

s = Solution()
print(s.countAndSay(6))
