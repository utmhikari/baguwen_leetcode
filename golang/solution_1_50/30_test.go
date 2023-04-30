package solution_1_50

import "fmt"

//30. 串联所有单词的子串
//给定一个字符串s和一些长度相同的单词words。
//找出 s 中恰好可以由words 中所有单词串联形成的子串的起始位置。
//
//注意子串要与words 中的单词完全匹配，中间不能有其他字符，
//但不需要考虑words中单词串联的顺序。




func findSubstring30(s string, words []string) []int {
	lw := len(words[0])
	ls := len(s)
	ans := make([]int, 0)
	if ls < lw {
		return ans
	}

	wordsCnt := make(map[string]int)
	counter := make(map[string]int)
	for _, w := range words {
		curCnt, ok := wordsCnt[w]
		if !ok {
			wordsCnt[w] = 1
			counter[w] = 0
		} else {
			wordsCnt[w] = curCnt + 1
		}
	}
	//fmt.Printf("wordsCnt: %+v\n", wordsCnt)

	// overall 0 ~ lw - 1 groups
	for i := 0; i < lw; i++ {
		left := i
		// reset counter
		for k, _ := range counter {
			counter[k] = 0
		}

		// cur index should always
		cur := left
		for cur <= ls - lw {
			subs := s[cur:cur + lw]
			curCnt, ok := counter[subs]
			if !ok {
				// one word mismatch, reset all counters
				left = cur + lw
				for k, _ := range counter {
					counter[k] = 0
				}
			} else {
				// add count to corresponding word
				counter[subs] = curCnt + 1
				// compare counters
				isNotEnough := false
				errWord := ""
				for w, c := range counter {
					wc, _ := wordsCnt[w]
					if c < wc {
						isNotEnough = true
					}
					if c > wc {
						errWord = w
						break
					}
				}

				fmt.Printf("[%d] counter: %+v\n", cur, counter)

				if !isNotEnough && len(errWord) == 0 {  // match all?
					ans = append(ans, left)
					oldLeftCnt, _ := counter[s[left:left+lw]]
					counter[s[left:left+lw]] = oldLeftCnt - 1
					left += lw
				} else if len(errWord) > 0 {
					newLeft := left
					for {
						// -1 counter for the word
						newSubs := s[newLeft:newLeft + lw]
						newLeftCnt, _ := counter[newSubs]
						counter[newSubs] = newLeftCnt - 1
						newLeft += lw
						left = newLeft

						// check if errWord got the correct counter
						curErrCnt, _ := counter[errWord]
						expectedErrCnt, _ := wordsCnt[errWord]
						if expectedErrCnt >= curErrCnt {
							break
						}
					}
				}
			}
			cur += lw
		}
	}

	return ans
}