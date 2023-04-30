import re


class Solution:
    def simplifyPath(self, path: str) -> str:
        p1 = re.sub(r'/+', '/', path).strip('/')
        paths = p1.split('/')
        ps = []
        for p in paths:
            if p != '..':
                if p != '.':
                    ps.append(p)
            elif len(ps) > 0:
                ps.pop()
        return '/' + '/'.join(ps)


s = Solution()
print(s.simplifyPath('/home/'))
print(s.simplifyPath('/../'))
print(s.simplifyPath('/home//foo/'))
print(s.simplifyPath('/a/./b/../../c/'))
print(s.simplifyPath('/a/../../b/../c//.//'))
print(s.simplifyPath('/a//b////c/d//././/..'))

