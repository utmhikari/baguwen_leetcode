package solution_101_150

import (
	"reflect"
	"testing"
)


var ans126 [][]string
var tmp126 []int
var visited126 map[int]int
var neighbor126 map[int]map[int]bool


func makeNeighbor126(i1 int, i2 int) {
	_, ok1 := neighbor126[i1]
	if !ok1 {
		neighbor126[i1] = make(map[int]bool)
	}
	_, ok2 := neighbor126[i2]
	if !ok2 {
		neighbor126[i2] = make(map[int]bool)
	}
	n1 := neighbor126[i1]
	n2 := neighbor126[i2]
	n1[i2] = true
	n2[i1] = true
}


func isNeighborStr126(s1 string, s2 string) bool {
	l := len(s1)
	if l != len(s2) || l == 0 {
		return false
	}

	diffCnt := 0
	for i := 0; i < l; i++ {
		if s1[i] != s2[i] {
			diffCnt++
		}
	}
	return diffCnt == 1
}


func isNeighbor126(i1 int, i2 int, wordList []string) bool {
	n, ok := neighbor126[i1]
	if ok {
		ok = n[i2]
		if ok {
			return true
		}
	}
	s1, s2 := wordList[i1], wordList[i2]
	return isNeighborStr126(s1, s2)
}


func bfs126(srcs []int, dest int, depth int) {
	nextSrcs := make([]int, 0)
	nextSrcMap := make(map[int]bool)
	for _, src := range srcs {
		n, ok := neighbor126[src]
		if !ok {
			continue
		}
		for i, _ := range n {
			d, _ := visited126[i]
			isRecorded, _ := nextSrcMap[i]
			if d == 0 && !isRecorded {
				visited126[i] = depth
				nextSrcMap[i] = true
				nextSrcs = append(nextSrcs, i)
			}
		}
	}

	//fmt.Printf("%+v, %d, %d\n", srcs, dest, depth)

	if len(nextSrcs) > 0 {
		bfs126(nextSrcs, dest, depth + 1)
	}
}


func backtrace126(endIdx int, wordList []string) {
	lastStep, _ := visited126[endIdx]
	if lastStep == 1 {
		ans126 = append(ans126, make([]string, len(tmp126)))
		for i := 0; i < len(tmp126); i++ {
			ans126[len(ans126) - 1][i] = wordList[tmp126[i]]
		}
	} else {
		n, _ := neighbor126[endIdx]
		for i, _ := range n {
			curStep, _ := visited126[i]
			if curStep == lastStep - 1 {
				tmp126 = append([]int{i}, tmp126...)
				backtrace126(i, wordList)
				tmp126 = tmp126[1:]
			}
		}
	}
}


func findLadders126(beginWord string, endWord string, wordList []string) [][]string {
	ans126 = make([][]string, 0)

	l := len(beginWord)
	if l == 0 || l != len(endWord) {
		return ans126
	}
	lenWords := len(wordList)
	if lenWords == 0 {
		return ans126
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
		return ans126
	}
	if lenWords == 1 {
		if isNeighborStr126(beginWord, endWord) {
			return [][]string{{endWord}}
		} else {
			return ans126
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
		return ans126
	}

	//fmt.Printf("%+v\n", visited126)

	tmp126 = []int{endIdx}
	backtrace126(endIdx, wordList)

	return ans126
}


type test126 struct {
	begin string
	end string
	words []string
	expected [][]string
}


func Test_126(t *testing.T) {
	inputs := []test126{
		{
			"hit",
			"cog",
			[]string{"hot","dot","dog","lot","log","cog"},
			[][]string{
				{"hit","hot","dot","dog","cog"},
				{"hit","hot","lot","log","cog"},
			},
		},
		{
			"hit",
			"cog",
			[]string{"hot","dot","dog","lot","log"},
			[][]string{},
		},
	}

	for i, input := range inputs {
		ans := findLadders126(input.begin, input.end, input.words)
		if !reflect.DeepEqual(ans, input.expected) {
			t.Errorf("failed at %d -> %+v, got %v\n", i, input, ans)
		} else {
			t.Logf("success at %d -> %+v\n", i, input)
		}
	}
}