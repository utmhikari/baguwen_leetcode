from typing import List


class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        l = len(intervals)
        if l <= 1:
            return intervals
        ans = list(intervals)
        tmp = []
        merged = True
        while merged:
            merged = False
            for item in ans:
                left, right = item[0], item[1]
                ltmp = len(tmp)
                isitemmerged = False
                for i in range(ltmp):
                    tmpitem = tmp[i]
                    tmpleft, tmpright = tmpitem[0], tmpitem[1]
                    if left <= tmpleft <= right or left <= tmpright <= right\
                            or (tmpleft < left and tmpright > right):
                        merged = True
                        isitemmerged = True
                        tmp[i] = [min(left, tmpleft), max(right, tmpright)]
                        break
                if not isitemmerged:
                    tmp.append([left, right])
            ans = tmp
            tmp = []
        return ans


s = Solution()
print(s.merge([[1,3]]))
print(s.merge([[1,3],[2,6],[8,10],[15,18]]))
print(s.merge([[1,4],[4,5]]))
print(s.merge([[1,4], [2,3]]))

