

class Solution:
    def evalRPN(self, tokens: [str]) -> int:
        stack = []
        for t in tokens:
            if t.isdigit() or (t.startswith('-') and t[1:].isdigit()):
                stack.append(t)
            else:
                a2 = int(stack.pop())
                a1 = int(stack.pop())
                print(t, a1, a2)
                if t == '+':
                    stack.append(a1 + a2)
                elif t == '-':
                    stack.append(a1 - a2)
                elif t == '*':
                    stack.append(a1 * a2)
                elif a1 * a2 > 0:
                    stack.append(a1 // a2)
                else:
                    stack.append(-(a1 // -a2))
        return stack[0] if len(stack) > 0 else 0

s = Solution()
print("-11".isnumeric())
print(s.evalRPN(["2", "1", "+", "3", "*"]))
print(s.evalRPN(["4", "13", "5", "/", "+"]))
print(s.evalRPN(["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]))