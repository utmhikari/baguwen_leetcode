package solution_101_150

import (
	"testing"
)

func ladderLength127(beginWord string, endWord string, wordList []string) int {
	l := len(beginWord)
	if l == 0 || l != len(endWord) {
		return 0
	}
	lenWords := len(wordList)
	if lenWords == 0 {
		return 0
	}
	neighbor126 = make(map[int]map[int]bool)
	endIdx := -1
	for i := 0; i < lenWords; i++ {
		if endWord == wordList[i] {
			endIdx = i
		}
		for j := i + 1 ; j < lenWords; j++ {
			if isNeighbor126(i, j, wordList) {
				makeNeighbor126(i, j)
			}
		}
	}
	if endIdx == -1 {
		return 0
	}
	if lenWords == 1 {
		if isNeighborStr126(beginWord, endWord) {
			return 2
		} else {
			return 0
		}
	}

	wordList = append(wordList, beginWord)
	for i := 0; i < lenWords; i++ {
		if isNeighbor126(i, lenWords, wordList) {
			makeNeighbor126(i, lenWords)
		}
	}

	visited126 = make(map[int]int)
	visited126[lenWords] = 1
	bfs126([]int{lenWords}, endIdx, 2)
	endVd := visited126[endIdx]
	if endVd == 0 {
		return 0
	} else {
		return endVd
	}
}



type test127 struct {
	begin string
	end string
	words []string
	expected int
}


func Test_127(t *testing.T) {
	inputs := []test127{
		{
			"hit",
			"cog",
			[]string{"hot","dot","dog","lot","log","cog"},
			5,
		},
		{
			"hit",
			"cog",
			[]string{"hot","dot","dog","lot","log"},
			0,
		},
	}

	for i, input := range inputs {
		ans := ladderLength127(input.begin, input.end, input.words)
		if ans != input.expected {
			t.Errorf("failed at %d -> %+v, got %v\n", i, input, ans)
		} else {
			t.Logf("success at %d -> %+v\n", i, input)
		}
	}
}