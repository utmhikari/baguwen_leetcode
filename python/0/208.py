class Trie:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.mp = dict()
        self.end = False

    def insert(self, word: str) -> None:
        """
        Inserts a word into the trie.
        """
        t = self
        for c in word:
            if c not in t.mp.keys():
                t.mp[c] = Trie()
            t = t.mp[c]
        t.end = True

    def search(self, word: str) -> bool:
        """
        Returns if the word is in the trie.
        """
        t = self
        for c in word:
            if c not in t.mp.keys():
                return False
            t = t.mp[c]
        return t.end

    def startsWith(self, prefix: str) -> bool:
        """
        Returns if there is any word in the trie that starts with the given prefix.
        """
        t = self
        for c in prefix:
            if c not in t.mp.keys():
                return False
            t = t.mp[c]
        return True