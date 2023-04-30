import json
import collections

def gettree(ps):
    t = {}
    for p in ps:
        _t = t
        for c in p:
            if c not in _t.keys():
                _t[c] = {}
            _t = _t[c]
        _t['end'] = True
    return t

def dfs(r, t, suf):
    if len(r) >= 3:
        return
    if 'end' in t.keys():
        r.append(suf)
        if len(r) >= 3:
            return
    for k in sorted(t.keys()):
        if k != 'end':
            suf += k
            dfs(r, t[k], suf)
            suf = suf[:-1]
            if len(r) >= 3:
                return

def getlist(t, c):
    if not t or c not in t.keys():
        return None, []
    r = []
    t = t[c]
    dfs(r, t, '')
    return t, r

class Solution:
    def suggestedProducts(self, products, searchWord):
        t = gettree(products)
        ans = []
        for i in range(len(searchWord)):
            c = searchWord[i]
            t, r = getlist(t, c)
            ans.append([searchWord[:i + 1] + w for w in r])
        return ans


s = Solution()

products = ["mobile","mouse","moneypot","monitor","mousepad"]
searchWord = "mouse"

print(json.dumps(s.suggestedProducts(products, searchWord), indent=2))

products = ["havana"]
searchWord = "havana"

print(json.dumps(s.suggestedProducts(products, searchWord), indent=2))

products = ["bags","baggage","banner","box","cloths"]
searchWord = "bags"

print(json.dumps(s.suggestedProducts(products, searchWord), indent=2))

products = ["havana"]
searchWord = "tatiana"
print(json.dumps(s.suggestedProducts(products, searchWord), indent=2))

