


def visit(s, p, l) -> int:
    l1, l2, r1, r2 = 0, 0, 0, 0
    add_min = 0
    mul_div = 0
    while p < l:
        if '0' <= s[p] <= '9':
            if add_min == 0:
                if mul_div == 0:
                    l1 = l1 * 10 + int(s[p])
                else:
                    l2 = l2 * 10 + int(s[p])
            elif mul_div == 0:
                r1 = r1 * 10 + int(s[p])
            else:
                r2 = r2 * 10 + int(s[p])
        elif s[p] in ['+', '-']:
            if add_min == 0:
                if mul_div != 0:
                    l1 = l1 * l2 if mul_div == 1 else l1 // l2
                    l2 = 0
            else:
                if mul_div != 0:
                    r1 = r1 * r2 if mul_div == 1 else r1 // r2
                    r2 = 0
                l1 = l1 + r1 if add_min == 1 else l1 - r1
                r1 = 0
            mul_div = 0
            add_min = 1 if s[p] == '+' else 2
        elif s[p] in ['*', '/']:
            if add_min == 0:
                if mul_div != 0:
                    l1 = l1 * l2 if mul_div == 1 else l1 // l2
                    l2 = 0
            else:
                if mul_div != 0:
                    r1 = r1 * r2 if mul_div == 1 else r1 // r2
                    r2 = 0
            mul_div = 1 if s[p] == '*' else 2
        p += 1
    if add_min == 0:
        if mul_div == 0:
            return l1
        else:
            return l1 * l2 if mul_div == 1 else l1 // l2
    elif mul_div == 0:
        return l1 + r1 if add_min == 1 else l1 - r1
    elif mul_div == 1:
        return l1 + (r1 * r2) if add_min == 1 else l1 - (r1 * r2)
    return l1 + (r1 // r2) if add_min == 1 else l1 - (r1 // r2)


class Solution:
    def calculate(self, s: str) -> int:
        cs = s.replace(' ', '')
        return visit(cs, 0, len(cs))

s = Solution()

print(s.calculate("3+2*2"))
print(s.calculate(" 3/2 "))
print(s.calculate(" 3+5 / 2 "))
print(s.calculate("1-1+1"))