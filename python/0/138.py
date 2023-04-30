from node import Node



class Solution:
    def copyRandomList(self, head: 'Node') -> 'Node':
        if not head:
            return
        l = list()
        d = dict()
        di = dict()
        p = head
        i = 0
        while p:
            di[p] = i
            l.append(Node(p.val, None, None))
            if p.random == p:
                l[-1].random = l[-1]
            elif p.random:
                if p.random not in d.keys():
                    d[p.random] = []
                d[p.random].append(i)
            if i > 0:
                l[-2].next = l[-1]
            p = p.next
            i += 1
        for k in di.keys():
            idx = di[k]
            if k in d.keys():
                for src in d[k]:
                    l[src].random = l[idx]
        return l[0]
