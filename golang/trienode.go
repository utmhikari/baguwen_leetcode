package main

type TrieNode struct {
	Dict map[string]*TrieNode
	IsEnd bool
}


func NewTrieNode() *TrieNode {
	return &TrieNode{
		Dict: make(map[string]*TrieNode),
		IsEnd: false,
	}
}


func (t *TrieNode) Insert(s string) {
	l := len(s)

	if l == 0 {
		return
	}

	p := t

	for i := 0; i < l; i++ {
		c := s[i:i+1]
		node, ok := p.Dict[c]
		if !ok {
			newNode := NewTrieNode()
			p.Dict[c] = newNode
			p = newNode
		} else {
			p = node
		}
	}

	p.IsEnd = true
}

func (t *TrieNode) InsertWords(words []string) {
	for _, word := range words {
		t.Insert(word)
	}
}

func (t *TrieNode) GetPrefixesByIdx(s string, left int) []string {
	prefixes := make([]string, 0)
	l := len(s)
	if l - left <= 0 {
		return prefixes
	}

	i, p := left, t
	for {
		if i >= l {
			return prefixes
		}

		c := s[i:i+1]
		node, ok := p.Dict[c]
		if !ok {
			return prefixes
		} else {
			if node.IsEnd {
				prefixes = append(prefixes, s[left:i+1])
			}
			p = node
		}

		i++
	}
}

func (t *TrieNode) GetPrefixIndicesByIdx(s string, left int) []int {
	prefixIndices := make([]int, 0)
	l := len(s)
	if l - left <= 0 {
		return prefixIndices
	}

	i, p := left, t
	for {
		if i >= l {
			break
		}

		c := s[i:i+1]
		node, ok := p.Dict[c]
		if !ok {
			break
		} else {
			if node.IsEnd {
				prefixIndices = append(prefixIndices, i + 1)
			}
			p = node
		}

		i++
	}

	lenIndices := len(prefixIndices)

	for x := 0; x < lenIndices / 2; x++ {
		prefixIndices[x], prefixIndices[lenIndices - 1 - x] = prefixIndices[lenIndices - 1 - x], prefixIndices[x]
	}

	return prefixIndices
}


func (t *TrieNode) GetPrefixes(s string) []string {
	return t.GetPrefixesByIdx(s, 0)
}
