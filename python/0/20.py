class Solution:
    def isValid(self, s: str) -> bool:
        l = [0]
        for c in s:
            if c == '(':
                l.append(1)
            elif c == '[':
                l.append(2)
            elif c == '{':
                l.append(3)
            elif c == ')':
                if l.pop() != 1:
                    return False
            elif c == ']':
                if l.pop() != 2:
                    return False
            elif c == '}':
                if l.pop() != 3:
                    return False
        return len(l) == 1
