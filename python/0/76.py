class Solution:
    def minWindow(self, s: str, t: str) -> str:
        d = {}
        leftidx = -1
        l = len(s)
        rightidx = l
        cnt = 0
        cntsum = {}
        for c in t:
            d[c] = 0
            if c not in cntsum.keys():
                cntsum[c] = 0
            cntsum[c] += 1
        lt = len(cntsum.keys())
        left = 0
        right = -1
        while True:
            right += 1
            if right == l:
                break
            c = s[right]
            if c not in d.keys():
                continue
            d[c] += 1
            if d[c] == cntsum[c]:
                # print(s, c, right, d[c], cntsum[c])
                cnt += 1
                if cnt == lt:
                    # find most left
                    while left <= right:
                        cl = s[left]
                        if cl not in d.keys():
                            left += 1
                            continue
                        if d[cl] > cntsum[cl]:
                            d[cl] -= 1
                        else:
                            d[cl] = cntsum[cl] - 1
                            cnt -= 1
                            # print('first: ', s[left:right + 1])
                            if right - left < rightidx - leftidx:
                                leftidx, rightidx = left, right
                            # find second start
                            left += 1
                            while left <= right:
                                cll = s[left]
                                if cll not in d.keys():
                                    left += 1
                                    continue
                                if d[cll] > cntsum[cll]:
                                    d[cll] -= 1
                                    left += 1
                                else:
                                    break
                            # print('second: ', s[left:right + 1])
                            break
                        left += 1
        if leftidx == -1:
            return ''
        return s[leftidx:rightidx + 1]


s = Solution()
print(s.minWindow("ADOBECODEBANC", "ABC"))
print(s.minWindow('a','a'))
print(s.minWindow('aa', 'aa'))
print(s.minWindow("adobecodebanc", "abcda"))



