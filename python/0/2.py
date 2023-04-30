# 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
#
# 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
#
# 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

from node import ListNode

class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        p1 = l1
        p2 = l2
        a = None
        t = None
        up = 0
        while True:
            if not p1 and not p2:
                if up == 1:
                    t.next = ListNode(1)
                break
            v1 = p1.val if p1 else 0
            v2 = p2.val if p2 else 0
            v = v1 + v2 + up
            if v >= 10:
                v -= 10
                up = 1
            else:
                up = 0
            if not a:
                a = ListNode(v)
                t = a
            else:
                t.next = ListNode(v)
                t = t.next
            if p1:
                p1 = p1.next
            if p2:
                p2 = p2.next
        return a

s = Solution()


