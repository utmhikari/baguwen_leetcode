class Solution:
    def isPalindrome(self, x: int) -> bool:
        s = str(x)
        l = 0
        r = len(s) - 1
        if r == 0:
            return True
        while l < r:
            if s[l] != s[r]:
                return False
            l += 1
            r -= 1
        return True