package solution_151_200



//187. 重复的DNA序列
//所有 DNA 都由一系列缩写为 'A'，'C'，'G' 和 'T' 的核苷酸组成，例如："ACGAATTCCG"。在研究 DNA 时，识别 DNA 中的重复序列有时会对研究非常有帮助。
//
//编写一个函数来找出所有目标子串，目标子串的长度为 10，且在 DNA 字符串 s 中出现次数超过一次。



type TrieNode187 struct {
	Dict map[string]*TrieNode187
	LeftIdx int
}


func NewTrieNode187() *TrieNode187 {
	return &TrieNode187{
		Dict: make(map[string]*TrieNode187),
		LeftIdx: -1,
	}
}


func (t *TrieNode187) Insert(s string, left int, right int) int {
	l := len(s)

	if l == 0 {
		return -1
	}

	p := t

	for i := left; i < right; i++ {
		c := s[i:i+1]
		node, ok := p.Dict[c]
		if !ok {
			newNode := NewTrieNode187()
			p.Dict[c] = newNode
			p = newNode
		} else {
			p = node
		}
	}

	if p.LeftIdx == -1 {
		p.LeftIdx = left
		return -1
	}
	return p.LeftIdx
}



func findRepeatedDnaSequences(s string) []string {
	l := len(s)
	b := 10
	if l < b {
		return make([]string, 0)
	}
	left := 0
	t := NewTrieNode187()
	cache := make(map[int]bool)
	for left + b <= l {
		leftIdx := t.Insert(s, left, left + b)
		if leftIdx != -1 {
			cache[leftIdx] = true
		}
		left++
	}
	ans := make([]string, len(cache))
	cnt := 0
	for leftIndex, _ := range cache {
		ans[cnt] = s[leftIndex:leftIndex+10]
		cnt++
	}
	return ans
}