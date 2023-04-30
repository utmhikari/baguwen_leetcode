class Solution:
    def longestPalindrome(self, s: str) -> str:
        l = len(s)
        if l == 0:
            return ""
        mx = 1
        mxstr = s[0]
        # single
        i = 0
        while True:
            if i >= l:
                break
            cur = 1
            left = i - 1
            right = i + 1
            while True:
                if left < 0 or right >= l or s[left] != s[right]:
                    if cur > mx:
                        mx = cur
                        mxstr = s[left + 1:right]
                    break
                else:
                    cur += 2
                    left -= 1
                    right += 1
            i += 1
        # double
        i = 0
        while True:
            if i + 1 >= l:
                break
            if s[i] == s[i + 1]:
                cur = 2
                left = i - 1
                right = i + 2
                while True:
                    if left < 0 or right >= l or s[left] != s[right]:
                        if cur > mx:
                            mx = cur
                            mxstr = s[left + 1:right]
                        break
                    else:
                        cur += 2
                        left -= 1
                        right += 1
            i += 1
        return mxstr


s = Solution()

a = "babad"
print(s.longestPalindrome(a))

a = "cbbd"
print(s.longestPalindrome(a))
