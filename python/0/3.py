# 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        l = len(s)
        if l == 0:
            return 0
        m = 0
        i = 1
        cs = dict()
        cs[s[0]] = 0
        start = 0
        while i < l:
            if s[i] in cs.keys():
                if i - start > m:
                    m = i - start
                newstart = cs[s[i]] + 1
                j = start
                while j < newstart:
                    del cs[s[j]]
                    j += 1
                start = newstart
            cs[s[i]] = i
            i += 1
        if i - start > m:
            m = i - start
        return m


s = Solution()

print(s.lengthOfLongestSubstring("abcabcbb"))

print(s.lengthOfLongestSubstring("bbbbb"))

print(s.lengthOfLongestSubstring("pwwkew"))

print(s.lengthOfLongestSubstring("abba"))

