package solution_51_100

import (
	"fmt"
	"strings"
)

//93. 复原 IP 地址
//给定一个只包含数字的字符串，用以表示一个 IP 地址，
//返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。
//
//有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），
//整数之间用 '.' 分隔。
//
//例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，
//但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。


var ans93 []string
var tmp93 []string

func solve93(s string, i int, pos int) {
	if pos > 3 {
		return
	}
	l := len(s)
	if i >= l {
		return
	}
	numLeftChars, numLeftPoses := l - i, 4 - pos
	if numLeftChars < numLeftPoses {
		return
	}

	sm := 0
	if pos == 3 {
		// has 0 as prefix?
		if s[i] == '0' {
			if i < l - 1 {
				return
			}
		} else {
			for j := i; j < l; j++ {
				digit := int(s[j] - '0')
				sm = sm * 10 + digit
				if sm > 255 {
					return
				}
			}
			if sm > 255 {
				return
			}
		}
		// join all
		tmp93[pos] = fmt.Sprintf("%d", sm)
		ans93 = append(ans93, strings.Join(tmp93, "."))
	} else {
		// 1 char
		tmp93[pos] = s[i:i+1]
		solve93(s, i + 1, pos + 1)
		sm = int(s[i] - '0')

		if s[i] != '0' {
			// 2 chars
			if i + 1 < l {
				sm = sm * 10 + int(s[i + 1] - '0')
				tmp93[pos] = s[i:i+2]
				solve93(s, i + 2, pos + 1)
				if i + 2 < l {
					// 3 chars
					sm = sm * 10 + int(s[i + 2] - '0')
					if sm <= 255 {
						tmp93[pos] = s[i:i+3]
						solve93(s, i + 3, pos + 1)
					}
				}
			}
		}
	}
}

func restoreIpAddresses93(s string) []string {
	ans93 = make([]string, 0)

	l := len(s)
	if l < 4 {
		return ans93
	}

	tmp93 = make([]string, 4)
	solve93(s, 0, 0)
	return ans93
}

