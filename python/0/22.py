import pprint
d = {1: {'()'},
     2: {'(())', '()()'},
     3: {'((()))', '()()()', '(())()', '(()())', '()(())'},
     4: {'(((())))',
         '((()()))',
         '((())())',
         '((()))()',
         '(()(()))',
         '(()()())',
         '(()())()',
         '(())(())',
         '(())()()',
         '()((()))',
         '()(()())',
         '()(())()',
         '()()(())',
         '()()()()'}}


def merge(st1, st2):
    st = set()
    if len(st1) == 1:
        for s in st2:
            st.add('(%s)' % s)
            st.add('()%s' % s)
            st.add('%s()' % s)
    else:
        for s1 in st1:
            for s2 in st2:
                st.add(s1 + s2)
                st.add(s2 + s1)
    return st

def gp(n):
    if n in d.keys():
        return d[n]
    st = set()
    mid = n // 2
    for i in range(mid):
        st |= merge(gp(i + 1), gp(n - i - 1))
    d[n] = st
    return st


class Solution:
    def generateParenthesis(self, n: int) -> [str]:
        return list(gp(n))

s = Solution()
a = s.generateParenthesis(8)
# print(a)
pprint.pprint(d)

