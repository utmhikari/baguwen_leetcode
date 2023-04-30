package solution_351_400

//336. 回文对
//
//给定一组 互不相同 的单词， 找出所有 不同 的索引对 (i, j)，
//使得列表中的两个单词， words[i] + words[j] ，可拼接成回文串。


type Node336 struct {
	ch [26]int
	flag int
}


var tree336 []Node336

func insert336(s string, id int) {
	add := 0
	for i := 0; i < len(s); i++ {
		x := int(s[i] - 'a')
		if tree336[add].ch[x] == 0 {
			tree336 = append(tree336, Node336{[26]int{}, -1})
			tree336[add].ch[x] = len(tree336) - 1
		}
		add = tree336[add].ch[x]
	}
	tree336[add].flag = id
}


func findWord336(s string, left int, right int) int {
	add := 0
	for i := right; i >= left; i-- {
		x := int(s[i] - 'a')
		if tree336[add].ch[x] == 0 {
			return -1
		}
		add = tree336[add].ch[x]
	}
	return tree336[add].flag
}

func isPalindrome336(s string, left int, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}


func palindromePairs(words []string) [][]int {
	tree336 = []Node336{{[26]int{}, -1}}
	n := len(words)
	for i := 0; i < n; i++ {
		insert336(words[i], i)
	}
	ret := make([][]int, 0)
	for i := 0; i < n; i++ {
		word := words[i]
		m := len(word)
		for j := 0; j <= m; j++ {
			if isPalindrome336(word, j, m - 1) {
				leftId := findWord336(word, 0, j - 1)
				if leftId != -1 && leftId != i {
					ret = append(ret, []int{i, leftId})
				}
			}
			if j != 0 && isPalindrome336(word, 0, j - 1) {
				rightId := findWord336(word, j, m - 1)
				if rightId != -1 && rightId != i {
					ret = append(ret, []int{rightId, i})
				}
			}
		}
	}
	return ret
}