from typing import List
import heapq

class Solution:
    def dailyTemperatures(self, T: List[int]) -> List[int]:
        l = len(T)
        if l == 0:
            return []
        lst = [0] * l
        if l == 1:
            return lst
        hp = []
        heapq.heappush(hp, (T[0], 0))
        i = 1
        while i < l:
            while i < l and T[i] > T[i - 1]:
                while len(hp) > 0:
                    item = hp[0]
                    if T[i] > item[0]:
                        lst[item[1]] = i - item[1]
                        heapq.heappop(hp)
                    else:
                        break
                lst[i - 1] = 1
                i += 1
            if i >= l:
                break
            heapq.heappush(hp, (T[i - 1], i - 1))
            i += 1
        return lst


s = Solution()
print(s.dailyTemperatures([73, 74, 75, 71, 69, 72, 76, 73]))