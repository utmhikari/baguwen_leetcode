import random

class RandomizedSet:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.keys = []
        self.d = dict()
        self.cd = dict()

    def insert(self, val: int) -> bool:
        """
        Inserts a value to the set. Returns true if the set did not already contain the specified element.
        """
        if val in self.d.keys():
            return False
        self.d[val] = len(self.keys)
        self.keys.append(val)
        return True

    def remove(self, val: int) -> bool:
        """
        Removes a value from the set. Returns true if the set contained the specified element.
        """
        if val not in self.d.keys():
            return False
        i = self.d[val]
        t = self.keys[-1]
        self.keys[i] = t
        self.d[t] = i
        del self.d[val]
        self.keys.pop()
        return True

    def getRandom(self) -> int:
        """
        Get a random element from the set.
        """
        i = random.randint(0, len(self.keys) - 1)
        return self.keys[i]

randomSet = RandomizedSet()

randomSet.insert(1)


randomSet.remove(2)


randomSet.insert(2)


randomSet.getRandom()


randomSet.remove(1)


randomSet.insert(2)


randomSet.getRandom()
