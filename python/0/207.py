class Node:
    def __init__(self, val):
        self.val = val
        self.prevs = set()
        self.nexts = set()


def findNexts(start, target, visited):
    if target in start:
        return True
    for st in start:
        if st not in visited:
            if findNexts(st.nexts, target, visited):
                return True
    return False

def findPrevs(start, target, visited):
    if target in start:
        return True
    for st in start:
        if st not in visited:
            if findPrevs(st.prevs, target, visited):
                return True
    return False

def isRound(nnext, nprev):
    nnext.prevs.add(nprev)
    nprev.nexts.add(nnext)
    if findNexts([nnext], nprev, set()):
        return True
    if findPrevs([nprev], nnext, set()):
        return True
    return False


class Solution:
    def canFinish(self, numCourses: int, prerequisites: [[int]]) -> bool:
        cs = dict()
        for pr in prerequisites:
            if not pr[0] in cs.keys():
                cs[pr[0]] = Node(pr[0])
            if not pr[1] in cs.keys():
                cs[pr[1]] = Node(pr[1])
            if isRound(cs[pr[0]], cs[pr[1]]):
                return False
        return True


s = Solution()
print(s.canFinish(2, [[1, 0]]))
print(s.canFinish(2, [[1,0],[0,1]]))
print(s.canFinish(1, []))


def canfinish(numCourses, prerequisites):
    indegrees = [0 for _ in range(numCourses)]
    adjacency = [[] for _ in range(numCourses)]
    queue = []
    # Get the indegree and adjacency of every course.
    for cur, pre in prerequisites:
        indegrees[cur] += 1
        adjacency[pre].append(cur)
    # Get all the courses with the indegree of 0.
    for i in range(len(indegrees)):
        if not indegrees[i]:
            queue.append(i)
    # BFS TopSort.
    while queue:
        pre = queue.pop(0)
        numCourses -= 1
        for cur in adjacency[pre]:
            indegrees[cur] -= 1
            if not indegrees[cur]:
                queue.append(cur)
    return not numCourses

