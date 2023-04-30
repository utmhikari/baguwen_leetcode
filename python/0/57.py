"""
给出一个无重叠的 ，按照区间起始端点排序的区间列表。

在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

示例 1:

输入: intervals = [[1,3],[6,9]], newInterval = [2,5]
输出: [[1,5],[6,9]]
示例 2:

输入: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出: [[1,2],[3,10],[12,16]]
解释: 这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
"""

from typing import List


class Solution:
    def insert(self,
               intervals: List[List[int]],
               newInterval: List[int]) -> List[List[int]]:
        l = len(intervals)
        if l == 0:
            return [newInterval]
        newIntervals = []
        if newInterval[1] < intervals[0][0]:
            newIntervals.append(newInterval)
            newIntervals.extend(intervals)
            return newIntervals
        if newInterval[0] > intervals[-1][1]:
            newIntervals.extend(intervals)
            newIntervals.append(newInterval)
            return newIntervals
        # find left pos and right pos
        left, right = -0.5, l - 0.5
        for i in range(l):
            if intervals[i][0] > newInterval[0]:
                # 不在上一个范围
                left = i - 0.5
                break
            if intervals[i][0] <= newInterval[0] <= intervals[i][1]:
                left = int(i)
                break
        for i in range(l):
            if intervals[l - 1 - i][1] < newInterval[1]:
                right = l - 1 - i + 0.5
                break
            if intervals[l - 1 - i][0] <= newInterval[1] <= intervals[l - 1 - i][1]:
                right = int(l - 1 - i)
                break

        left_pos = left if isinstance(left, int) else int(left + 0.5)
        right_pos = right if isinstance(right, int) else int(right - 0.5)
        print(left_pos, right_pos)
        # 几种情况
        for i in range(left_pos):
            newIntervals.append(intervals[i])
        newIntervals.append([
            intervals[left_pos][0] if isinstance(left, int) else newInterval[0],
            intervals[right_pos][1] if isinstance(right, int) else newInterval[1]
        ])
        for i in range(right_pos + 1, l):
            newIntervals.append(intervals[i])
        return newIntervals


if __name__ == '__main__':
    s = Solution()
    intervals = [[1,3],[6,9]]
    newInterval = [2,5]
    print(s.insert(intervals, newInterval))
    intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]]
    newInterval = [4,8]
    print(s.insert(intervals, newInterval))
    intervals = [[1,5]]
    newInterval = [2,7]
    print(s.insert(intervals, newInterval))
    intervals = [[1,2]]
    newInterval = [-1,0]
    print(s.insert(intervals, newInterval))
    intervals = [[1,5]]
    newInterval = [6,7]
    print(s.insert(intervals, newInterval))