package solution_201_250



//211. 添加与搜索单词 - 数据结构设计


type WordDictionary struct {
	Chars map[rune]*WordDictionary
	IsEnd bool
}


/** Initialize your data structure here. */
func Constructor211() WordDictionary {
	return WordDictionary{
		Chars: make(map[rune]*WordDictionary),
		IsEnd: false,
	}
}


func (this *WordDictionary) AddWord(word string)  {
	p := this
	for _, r := range []rune(word) {
		nxt, ok := p.Chars[r]
		if !ok {
			nxtNode := Constructor211()
			p.Chars[r] = &nxtNode
			p = &nxtNode
		} else {
			p = nxt
		}
	}
	p.IsEnd = true
}


func (this *WordDictionary) SearchByRune(word string, idx int) bool {
	if idx == len(word) {
		return this.IsEnd
	}
	r := rune(word[idx])
	if r == '.' {
		for _, nxtwd := range this.Chars {
			if nxtwd.SearchByRune(word, idx + 1) {
				return true
			}
		}
	} else {
		nxt, ok := this.Chars[r]
		if ok {
			return nxt.SearchByRune(word, idx + 1)
		}
	}
	return false
}


func (this *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		return false
	}
	return this.SearchByRune(word, 0)
}
