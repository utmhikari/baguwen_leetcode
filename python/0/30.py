from typing import List
import copy


def fs(s, start, end, wordsdict, lw, curcnt, wordscnt, ans):
    if start >= end:
        return
    ss = s[start:start + lw]
    print(s, start, end, curcnt, wordscnt, ss)
    if ss in wordsdict.keys() and wordsdict[ss] > 0:
        wordsdict[ss] -= 1
        curcnt += 1
        if curcnt == wordscnt:
            curcnt -= 1
            ans.append(start - curcnt * lw)
            first_s = s[start - curcnt * lw:start - (curcnt - 1) * lw]
            wordsdict[first_s] += 1
    elif ss in wordsdict.keys():
        for i in range(curcnt):
            prev_s = s[start - curcnt * lw:start - (curcnt - 1) * lw]
            wordsdict[prev_s] += 1
            curcnt -= 1
            if prev_s == ss:
                wordsdict[prev_s] -= 1
                curcnt += 1
                break
    else:
        for i in range(curcnt):
            wordsdict[s[start - (i + 1) * lw:start - i * lw]] += 1
        curcnt = 0
    fs(s, start + lw, end, wordsdict, lw, curcnt, wordscnt, ans)


class Solution:
    def findSubstring(self, s: str, words: List[str]) -> List[int]:
        l = len(s)
        if len(words) == 0:
            return []
        lw = len(words[0])
        wordslen = lw * len(words)
        ans = []
        wordsdict = {}
        for w in words:
            if w not in wordsdict.keys():
                wordsdict[w] = 0
            wordsdict[w] += 1
        for i in range(lw):
            start = i
            end = i + ((l - 1 - i) // lw * lw)
            if end + lw == l:
                end = l
            if end - start < wordslen:
                break
            wd = copy.deepcopy(wordsdict)
            fs(s, start, end, wd, lw, 0, len(words), ans)
        return ans




s = Solution()
print(s.findSubstring("barfoothefoobarman", ["foo","bar"]))
print(s.findSubstring("wordgoodgoodgoodbestword", ["word","good","best","word"]))
print(s.findSubstring("wordgoodgoodgoodbestword", ["word","good","best","good"]))
print(s.findSubstring("aaaaaaaa", ["aa","aa","aa"]))