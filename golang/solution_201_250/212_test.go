package solution_201_250

import "bytes"

//212. 单词搜索 II
//给定一个m x n 二维字符网格board和一个单词（字符串）列表 words，
//找出所有同时在二维网格和字典中出现的单词。
//
//单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，
//其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
//同一个单元格内的字母在一个单词中不允许被重复使用。


type Trie212 struct {
	Chars map[rune]*Trie212
	IsEnd bool
	IsMarked bool
}


/** Initialize your data structure here. */
func Constructor212() Trie212 {
	return Trie212{
		Chars: make(map[rune]*Trie212),
		IsEnd: false,
		IsMarked: false,
	}
}


/** Inserts a word into the trie. */
func (this *Trie212) Insert(word string)  {
	p := this
	for _, r := range []rune(word) {
		nxt, ok := p.Chars[r]
		if !ok {
			nxtNode := Constructor212()
			p.Chars[r] = &nxtNode
			p = &nxtNode
		} else {
			p = nxt
		}
	}
	p.IsEnd = true
}


var ans212 []string
var visited212 map[int]bool
var buffer212 bytes.Buffer
var trie212 Trie212

func find212(board [][]byte, i int, j int, trie *Trie212) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return
	}

	key := i * len(board[0]) + j
	_, ok := visited212[key]
	if ok {
		return
	}

	visited212[key] = true

	r := rune(board[i][j])
	nxtTrie, ok := trie.Chars[r]
	if ok {
		buffer212.WriteRune(r)

		if nxtTrie.IsEnd && !nxtTrie.IsMarked {
			nxtTrie.IsMarked = true
			ans212 = append(ans212, buffer212.String())
		}

		find212(board, i + 1, j, nxtTrie)
		find212(board, i - 1, j, nxtTrie)
		find212(board, i, j + 1, nxtTrie)
		find212(board, i, j - 1, nxtTrie)

		buffer212.Truncate(buffer212.Len() - 1)
	}

	delete(visited212, key)
}


func findWords212(board [][]byte, words []string) []string {
	ans212 = make([]string, 0)
	trie212 = Constructor212()
	for _, word := range words {
		trie212.Insert(word)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			visited212 = make(map[int]bool)
			buffer212.Reset()
			find212(board, i, j, &trie212)
		}
	}

	return ans212
}
