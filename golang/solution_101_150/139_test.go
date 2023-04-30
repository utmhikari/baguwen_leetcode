package solution_101_150

import (
	"fmt"
	"github.com/utmhikari/go-leetcode"
	"strconv"
	"testing"
)

//139. 单词拆分
//给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，
//判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。


func solve139(s string, left int, trieNode *main.TrieNode, cache *[]bool) bool {
	l := len(s)
	if left == l {
		return true
	}

	if (*cache)[left] {
		return false
	}

	prefixIndices := trieNode.GetPrefixIndicesByIdx(s, left)
	if len(prefixIndices) == 0 {
		(*cache)[left] = true
		return false
	}

	//fmt.Printf("%+v\n", prefixIndices)

	for _, prefixIdx := range prefixIndices {
		if solve139(s, prefixIdx, trieNode, cache) {
			return true
		}
	}

	(*cache)[left] = true
	return false
}


func wordBreak139(s string, wordDict []string) bool {
	if len(s) == 0 {
		return false
	}

	trieNode := main.NewTrieNode()
	trieNode.InsertWords(wordDict)

	cache := make([]bool, len(s))
	return solve139(s, 0, trieNode, &cache)
}


type test139 struct {
	s string
	wordDict []string
	expected bool
}


func Test_139 (t *testing.T) {
	cases := append(make([]test139, 0),
		test139{
			s:     "leetcode",
			wordDict:   []string{"leet", "code"},
			expected: true,
		},
		test139{
			s:     "applepenapple",
			wordDict:   []string{"apple", "pen"},
			expected: true,
		},
		test139{
			s:     "catsandog",
			wordDict:   []string{"cats", "dog", "sand", "and", "cat"},
			expected: false,
		},
		test139{
			s:     "a",
			wordDict:   []string{"a"},
			expected: true,
		},
		test139{
			s: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
			wordDict: []string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"},
			expected: false,
		},
	)

	var ans bool

	for _, c := range cases {
		ans = wordBreak139(c.s, c.wordDict)
		if ans != c.expected {
			t.Errorf("failed at %v+, ans: %s\n", c, strconv.FormatBool(ans))
		} else {
			fmt.Printf("success at %v+\n", c)
		}
	}
}