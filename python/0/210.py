class Solution:
    def findOrder(self, numCourses: int, prerequisites: [[int]]) -> [int]:
        indegrees = [0 for _ in range(numCourses)]
        nexts = [[] for _ in range(numCourses)]
        for n, p in prerequisites:
            indegrees[n] += 1
            nexts[p].append(n)
        queue = []
        for i in range(numCourses):
            if indegrees[i] == 0:
                queue.append(i)
        ans = []
        while queue:
            p = queue.pop(0)
            ans.append(p)
            numCourses -= 1
            for n in nexts[p]:
                indegrees[n] -= 1
                if indegrees[n] < 0:
                    return []
                if not indegrees[n]:
                    queue.append(n)
        return ans if not numCourses else []

s = Solution()
print(s.findOrder(4, [[1,0],[2,0],[3,1],[3,2]]))



