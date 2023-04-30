package solution_101_150

import (
	"fmt"
	"github.com/utmhikari/go-leetcode"
	"reflect"
	"strings"
	"testing"
)


var ansCache140 map[int][][]string


func solve140(s string, left int, trieNode *main.TrieNode, cache *[]bool) ([][]string, bool) {
	//fmt.Printf("<start> %s, %d, %+v\n", s, left, cache)

	if (*cache)[left] {
		return nil, false
	}

	curAns, ok := ansCache140[left]
	if ok {
		return curAns, true
	}

	prefixIndices := trieNode.GetPrefixIndicesByIdx(s, left)
	if len(prefixIndices) == 0 {
		(*cache)[left] = true
		return nil, false
	}

	ok = false
	curAns = make([][]string, 0)

	for _, prefixIdx := range prefixIndices {
		if prefixIdx == len(s) {
			curAns = append(curAns, []string{s[left:prefixIdx]})
			ok = true
		} else {
			nextAns, nextOk := solve140(s, prefixIdx, trieNode, cache)
			if nextOk {
				ok = true
				for _, an := range nextAns {
					newAns := make([]string, len(an) + 1)
					newAns[0] = s[left:prefixIdx]
					for i, s2 := range an {
						newAns[i+1] = s2
					}
					curAns = append(curAns, newAns)
				}
			}
		}
	}

	//fmt.Printf("<end> %s, %d, %+v, %+v\n", s, left, prefixIndices, curAns)

	if ok {
		ansCache140[left] = curAns
		return curAns, true
	} else {
		(*cache)[left] = true
		return nil, false
	}
}


func wordBreak140(s string, wordDict []string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}

	trieNode := main.NewTrieNode()
	trieNode.InsertWords(wordDict)

	cache := make([]bool, len(s))
	ansCache140 = make(map[int][][]string)

	solutions, ok := solve140(s, 0, trieNode, &cache)
	if !ok {
		return make([]string, 0)
	}

	ans := make([]string, len(solutions))
	for i, solution := range solutions {
		ans[i] = strings.Join(solution, " ")
	}
	return ans
}


type test140 struct {
	s string
	wordDict []string
	expected []string
}


func Test_140 (t *testing.T) {
	cases := append(make([]test140, 0),
		test140{
			s:     "catsanddog",
			wordDict:   []string{"cat", "cats", "and", "sand", "dog"},
			expected: []string{"cats and dog", "cat sand dog"},
		},
		test140{
			s:     "pineapplepenapple",
			wordDict:   []string{"apple", "pen", "applepen", "pine", "pineapple"},
			expected: []string{
				"pineapple pen apple",
				"pine applepen apple",
				"pine apple pen apple",
			},
		},
		test140{
			s:     "a",
			wordDict:   []string{"a"},
			expected: []string{"a"},
		},
	)

	var ans []string

	for _, c := range cases {
		ans = wordBreak140(c.s, c.wordDict)
		if !reflect.DeepEqual(ans, c.expected) {
			t.Errorf("failed at %+v, ans: %+v\n", c, ans)
		} else {
			fmt.Printf("success at %+v\n", c)
		}
	}
}
