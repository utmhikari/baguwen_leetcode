package solution_351_400

import "sort"

//354. 俄罗斯套娃信封问题
//
//给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。
//
//当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
//
//请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。



func maxEnvelopes(envelopes [][]int) int {
	l := len(envelopes)
	if l <= 1 {
		return l
	}

	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	ans := make([]int, l)
	for i := 0; i < l; i++ {
		ans[i] = 1
	}

	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if envelopes[j][1] < envelopes[i][1] {
				if ans[j] + 1 > ans[i] {
					ans[i] = ans[j] + 1
				}
			}
		}
	}

	mx := 1
	for i := 0; i < l; i++ {
		if ans[i] > mx {
			mx = ans[i]
		}
	}
	return mx
}
