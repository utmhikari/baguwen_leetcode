class Solution:
    def fizzBuzz(self, n: int) -> [str]:
        a = []
        for i in range(n):
            j = i + 1
            if j % 3 == 0 and j % 5 == 0:
                a.append("FizzBuzz")
            elif j % 3 == 0:
                a.append("Fizz")
            elif j % 5 == 0:
                a.append("Buzz")
            else:
                a.append(str(j))
        return a