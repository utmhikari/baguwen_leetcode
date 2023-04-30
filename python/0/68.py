from typing import List


def justify(tmp, maxWidth, last=False):
    if last:
        s = ' '.join(tmp)
        return s + (maxWidth - len(s)) * ' '
    l = len(tmp)
    if l == 1:
        return tmp[0] + (maxWidth - len(tmp[0])) * ' '
    wl = 0
    for w in tmp:
        wl += len(w)
    d = (maxWidth - wl) // (l - 1)
    md = maxWidth - wl - d * (l - 1)
    sizes = [d] * (l - 1)
    for i in range(md):
        sizes[i] += 1
    s = ''
    for i in range(l - 1):
        s += tmp[i]
        s += sizes[i] * ' '
    s += tmp[-1]
    return s


class Solution:
    def fullJustify(self, words: List[str], maxWidth: int) -> List[str]:
        ans = []
        l = len(words)
        i = 0
        tmp = []
        sm = 0
        while i < l:
            w = words[i]
            lw = len(w)
            if len(tmp) == 0:
                tmp.append(w)
                sm += lw
            else:
                if lw + 1 + sm <= maxWidth:
                    tmp.append(w)
                    sm = sm + lw + 1
                else:
                    ans.append(justify(tmp, maxWidth))
                    tmp = []
                    sm = 0
                    continue
            if i == l - 1:
                ans.append(justify(tmp, maxWidth, last=True))
                return ans
            i += 1
        return ans

s = Solution()
print(s.fullJustify(["This", "is", "an", "example", "of", "text", "justification."], 16))
print(s.fullJustify(["What","must","be","acknowledgment","shall","be"], 16))
print(s.fullJustify(["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"], 20))
