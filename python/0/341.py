from collections import deque

class NestedIterator:
    def __init__(self, nestedList):
        self.deque = deque()
        for nestedInteger in nestedList:
            self.deque.appendleft(nestedInteger)
        self.tmpNestedIterator = None

    def next(self) -> int:
        if self.tmpNestedIterator:
            return self.tmpNestedIterator.next()
        item = self.deque[-1]
        if item.isInteger():
            return self.deque.pop().getInteger()
        if not self.tmpNestedIterator:
            self.tmpNestedIterator = NestedIterator(item.getList())
        return self.next()

    def hasNext(self) -> bool:
        if self.tmpNestedIterator:
            if not self.tmpNestedIterator.hasNext():
                self.tmpNestedIterator = None
                self.deque.pop()
            else:
                return True
        if len(self.deque) == 0:
            return False
        item = self.deque[-1]
        if item.isInteger():
            return True
        self.tmpNestedIterator = NestedIterator(item.getList())
        return self.hasNext()



s = NestedIterator([[1,1],2,[1,1]])
a = []
while s.hasNext():
    a.append(s.next())
print(a)

s = NestedIterator([1,[4,[6]]])
a = []
while s.hasNext():
    a.append(s.next())
print(a)




