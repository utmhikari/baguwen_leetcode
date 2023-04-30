from node import ListNode


def mergeTwoLists(l1: ListNode, l2: ListNode) -> ListNode:
    p1 = l1
    p2 = l2
    first = ListNode(0)
    p = first
    while True:
        if not p1:
            p.next = p2
            break
        elif not p2:
            p.next = p1
            break
        elif p1.val <= p2.val:
            p.next = ListNode(p1.val)
            p1 = p1.next
        else:
            p.next = ListNode(p2.val)
            p2 = p2.next
        p = p.next
    return first.next


def mkl(lists, l, r):
    if r == l:
        return lists[l]
    elif r - l == 1:
        return mergeTwoLists(lists[l], lists[r])
    mid = l + (r - l) // 2
    return mergeTwoLists(mkl(lists, l, mid - 1), mkl(lists, mid, r))


class Solution:
    def mergeKLists(self, lists) -> ListNode:
        l = len(lists)
        if l != 0:
            return mkl(lists, 0, l - 1)

