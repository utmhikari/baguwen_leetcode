from node import ListNode

class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> ListNode:
        if not headA or not headB:
            return
        pa = headA
        pb = headB
        if pa == pb:
            return pa
        nodes = set()
        nodes.add(pa)
        nodes.add(pb)
        turn = 0
        while pa or pb:
            if not pb or turn == 0:
                if not pa:
                    break
                pa = pa.next
                if pb:
                    turn = 1
                if pa in nodes:
                    return pa
                elif pa:
                    nodes.add(pa)
            elif not pa or turn == 1:
                if not pb:
                    break
                pb = pb.next
                if pa:
                    turn = 0
                if pb in nodes:
                    return pb
                elif pb:
                    nodes.add(pb)
        return
