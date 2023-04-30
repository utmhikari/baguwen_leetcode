import pprint


class Solution:
    def canCompleteCircuit(self, gas: [int], cost: [int]) -> int:
        sgas = sum(gas)
        scost = sum(cost)
        if sgas < scost:
            return -1
        l = len(gas)
        pos_starts = dict()
        neg_starts = dict()
        neq_idxs, eq_idxs = list(), set()
        pos_start, neg_start, pos_sum, neg_sum = -1, -1, 0, 0
        for i in range(l):
            if gas[i] == cost[i]:
                eq_idxs.add(i)
            else:
                neq_idxs.append(i)
        ln = len(neq_idxs)
        if ln == 0:
            return 0
        for i in neq_idxs:
            c = gas[i] - cost[i]
            if c == 0:
                eq_idxs.add(i)
            elif c > 0:
                pos_sum += c
                if pos_start == -1:
                    pos_start = i
                    if neg_start != -1:
                        neg_starts[neg_start] = {'end': i - 1, 'sum': neg_sum}
                        neg_sum = 0
                        neg_start = -1
            else:
                neg_sum += c
                if neg_start == -1:
                    neg_start = i
                    if pos_start != -1:
                        pos_starts[pos_start] = {'end': i - 1, 'sum': pos_sum}
                        pos_sum = 0
                        pos_start = -1
            if i == neq_idxs[ln - 1]:
                j = 0
                while j < l and j in eq_idxs:
                    j += 1
                if j == l:
                    return 0
                if pos_start >= 0:
                    if j in pos_starts.keys():
                        pos_starts[pos_start] = {
                            'end': pos_starts[j]['end'],
                            'sum': pos_starts[j]['sum'] + pos_sum
                        }
                        del pos_starts[j]
                    else:
                        pos_starts[pos_start] = {
                            'end': l - 1,
                            'sum': pos_sum
                        }
                elif neg_start >= 0:
                    if j in neg_starts.keys():
                        neg_starts[neg_start] = {
                            'end': neg_starts[j]['end'],
                            'sum': neg_starts[j]['sum'] + neg_sum
                        }
                        del neg_starts[j]
                    else:
                        neg_starts[neg_start] = {
                            'end': l - 1,
                            'sum': neg_sum
                        }
        for k in pos_starts.keys():
            n = 0
            t = k
            while True:
                if t in pos_starts.keys():
                    end = pos_starts[t]['end']
                    n += pos_starts[t]['sum']
                    t = end + 1
                elif t in neg_starts.keys():
                    end = neg_starts[t]['end']
                    n += neg_starts[t]['sum']
                    if n < 0:
                        break
                    t = end + 1
                else:
                    t += 1
                t %= l
                if t == k:
                    return k
        return -1


s = Solution()
print(s.canCompleteCircuit([1,2,3,4,5], [3,4,5,1,2]))
print(s.canCompleteCircuit([2,3,4], [3,4,3]))
print(s.canCompleteCircuit([5,1,2,3,4],[4,4,1,5,1]))