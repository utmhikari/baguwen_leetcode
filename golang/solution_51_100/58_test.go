package solution_51_100

//58. 最后一个单词的长度
//给你一个字符串 s，由若干单词组成，单词之间用空格隔开。
//返回字符串中最后一个单词的长度。如果不存在最后一个单词，
//请返回 0
//
//单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。


func lengthOfLastWord58(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	leftIdx, rightIdx := -1, -1
	for i := l - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if rightIdx != -1 {
				leftIdx = i
				break
			}
		} else {
			if rightIdx == -1 {
				rightIdx = i
			}
		}
	}
	if rightIdx == -1 {
		return 0
	}
	return rightIdx - leftIdx
}