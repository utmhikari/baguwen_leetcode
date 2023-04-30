import pprint

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


def visit(r: TreeNode, a: [], d: int):
    if r:
        if len(a) <= d:
            a.append([])
        a[d].append(r.val)
        visit(r.left, a, d + 1)
        visit(r.right, a, d + 1)


def printtree(root: TreeNode) -> [[int]]:
    a = []
    visit(root, a, 0)
    pprint.pprint(a, indent=2)


class Node:
    def __init__(self, val, next, random):
        self.val = val
        self.next = next
        self.random = random
