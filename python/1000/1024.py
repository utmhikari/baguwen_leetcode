def vs(clips, visited, deleted, left, right, cur):
    # print("lr:", left, right)
    l = len(clips)
    leftidx, rightidx = -1, -1
    for i in range(l):
        if i in visited or i in deleted:
            continue
        if clips[i][0] <= left and clips[i][1] >= right:
            return cur + 1
        if clips[i][0] <= left:
            if leftidx == -1:
                leftidx = i
            elif clips[i][1] >= clips[leftidx][1]:
                deleted.add(leftidx)
                leftidx = i
            else:
                deleted.add(i)
        elif clips[i][1] >= right:
            if rightidx == -1:
                rightidx = i
            elif clips[i][0] <= clips[rightidx][0]:
                deleted.add(rightidx)
                rightidx = i
            else:
                deleted.add(i)
    if leftidx == -1 or rightidx == -1:
        return -1
    elif clips[leftidx][1] >= clips[rightidx][0]:
        return cur + 2
    visited.add(leftidx)
    visited.add(rightidx)
    return vs(clips, visited, deleted, clips[leftidx][1], clips[rightidx][0], cur + 2)


class Solution:
    def videoStitching(self, clips, T) -> int:
        visited = set()
        deleted = set()
        l = len(clips)
        for i in range(l):
            if clips[i][0] > T:
                deleted.add(i)
        ans = vs(clips, visited, deleted, 0, T, 0)
        return ans if ans > 0 else -1


s = Solution()

print(s.videoStitching([[0,4],[2,8]], 5))
print(s.videoStitching([[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]], 10))
print(s.videoStitching([[0,1],[1,2]], 5))
print(s.videoStitching([[0,2],[3,6], [7, 10]], 10))
print(s.videoStitching([[17,18],[25,26],[16,21],[3,3],[19,23],[1,5],[0,2],[9,20],[5,17],[8,10]], 15))