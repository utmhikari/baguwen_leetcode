package solution_351_400

import "strings"

//388. 文件的最长绝对路径


var ans388 int

func lengthLongestPath(input string) int {
	ans388 = 0
	lines := strings.Split(input, "\n")
	lvlLens := make([]int, 0)
	lvlLenSums := make([]int, 0)
	for _, line := range lines {
		lvl := 0
		for i, r := range line {
			if r != '\t' {
				lvl = i
				break
			}
		}
		f := line[lvl:]
		for len(lvlLens) <= lvl {
			lvlLens = append(lvlLens, 0)
			lvlLenSums = append(lvlLenSums, 0)
		}
		lvlLens[lvl] = len(f)
		if lvl == 0 {
			lvlLenSums[lvl] = len(f)
		} else {
			lvlLenSums[lvl] = len(f) + lvlLenSums[lvl-1] + 1  // separator + 1
		}
		if strings.Contains(f, ".") {
			if lvlLenSums[lvl] > ans388 {
				ans388 = lvlLenSums[lvl]
			}
		}
	}
	return ans388
}