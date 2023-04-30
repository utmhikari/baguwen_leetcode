package solution_51_100

import (
	"bytes"
	"fmt"
)

//68. 文本左右对齐
//给定一个单词数组和一个长度maxWidth，重新排版单词，
//使其成为每行恰好有maxWidth个字符，且左右两端对齐的文本。
//
//你应该使用“贪心算法”来放置给定的单词；也就是说，尽可能多地往每行中放置单词。
//必要时可用空格' '填充，使得每行恰好有 maxWidth个字符。
//
//要求尽可能均匀分配单词间的空格数量。
//如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
//
//文本的最后一行应为左对齐，且单词之间不插入额外的空格。



func fullJustify68(words []string, maxWidth int) []string {
	ans := make([]string, 0)
	l := len(words)
	if l == 0 {
		return ans
	}
	if l == 1 {
		var b bytes.Buffer
		b.WriteString(words[0])
		for i := len(words[0]); i < maxWidth; i++ {
			b.WriteRune(' ')
		}
		ans = append(ans, b.String())
		return ans
	}

	pl, pr, cl := 0, 1, len(words[0])
	for {
		if pr >= l || cl + len(words[pr]) + pr - pl > maxWidth {
			// handle pl ~ pr - 1
			var b bytes.Buffer
			if pr - 1 == pl {
				b.WriteString(words[pl])
				for i := len(words[pl]); i < maxWidth; i++ {
					b.WriteRune(' ')
				}
			} else if pr < l {
				numLeftSpaces := maxWidth - cl
				numWords := pr - pl
				gapNumSpaces := make([]int, numWords - 1)
				numAvgSpaces := numLeftSpaces / (numWords - 1)
				numModSpaces := numLeftSpaces % (numWords - 1)
				for i := 0; i < numWords - 1; i++ {
					gapNumSpaces[i] += numAvgSpaces
					if i < numModSpaces {
						gapNumSpaces[i]++
					}
				}
				fmt.Printf("cl: %d, gaps: %v\n", cl, gapNumSpaces)
				b.WriteString(words[pl])
				for i := pl + 1; i < pr; i++ {
					for j := 0; j < gapNumSpaces[i - pl - 1]; j++ {
						b.WriteRune(' ')
					}
					b.WriteString(words[i])
				}
			} else {
				b.WriteString(words[pl])
				for i := pl + 1; i < pr; i++ {
					b.WriteRune(' ')
					b.WriteString(words[i])
				}
				for i := b.Len(); i < maxWidth; i++ {
					b.WriteRune(' ')
				}
			}
			ans = append(ans, b.String())

			// switch to next
			pl = pr
			if pr >= l {
				break
			}
			cl = len(words[pr])
		} else {
			cl += len(words[pr])
		}
		pr++
	}

	return ans
}