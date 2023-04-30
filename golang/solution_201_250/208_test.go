package solution_201_250


//208 trie


type Trie struct {
	Chars map[rune]*Trie
	IsEnd bool
}


/** Initialize your data structure here. */
func Constructor208() Trie {
	return Trie{
		Chars: make(map[rune]*Trie),
		IsEnd: false,
	}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	p := this
	for _, r := range []rune(word) {
		nxt, ok := p.Chars[r]
		if !ok {
			nxtNode := Constructor208()
			p.Chars[r] = &nxtNode
			p = &nxtNode
		} else {
			p = nxt
		}
	}
	p.IsEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	p := this
	for _, r := range []rune(word) {
		nxt, ok := p.Chars[r]
		if !ok {
			return false
		}
		p = nxt
	}
	return p.IsEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	p := this
	for _, r := range []rune(prefix) {
		nxt, ok := p.Chars[r]
		if !ok {
			return false
		}
		p = nxt
	}
	return true
}
