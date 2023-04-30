package solution_151_200

import (
	"strconv"
	"strings"
)

//
//165. 比较版本号
//给你两个版本号 version1 和 version2 ，请你比较它们。
//
//版本号由一个或多个修订号组成，各修订号由一个 '.' 连接。每个修订号由 多位数字 组成，可能包含 前导零 。每个版本号至少包含一个字符。修订号从左到右编号，下标从 0 开始，最左边的修订号下标为 0 ，下一个修订号下标为 1 ，以此类推。例如，2.5.33 和 0.1 都是有效的版本号。
//
//比较版本号时，请按从左到右的顺序依次比较它们的修订号。比较修订号时，只需比较 忽略任何前导零后的整数值 。也就是说，修订号 1 和修订号 001 相等 。如果版本号没有指定某个下标处的修订号，则该修订号视为 0 。例如，版本 1.0 小于版本 1.1 ，因为它们下标为 0 的修订号相同，而下标为 1 的修订号分别为 0 和 1 ，0 < 1 。
//
//返回规则如下：
//
//如果version1>version2返回1，
//如果version1<version2 返回 -1，
//除此之外返回 0。


func getVersions165(v string) []int64 {
	sp := strings.Split(v, ".")
	intV := make([]int64, len(sp))
	for i := 0; i < len(sp); i++ {
		s := sp[i]
		j := 0
		for ; j < len(s); j++ {
			if s[j:j+1] != "0" {
				break
			}
		}
		if j == len(s) {
			intV[i] = 0
		} else {
			sv, _ := strconv.ParseInt(s[j:], 10, 64)
			intV[i] = sv
		}
	}
	return intV
}


func compareVersion165(version1 string, version2 string) int {
	sp1, sp2 := getVersions165(version1), getVersions165(version2)

	if len(sp1) < len(sp2) {
		sp1 = append(sp1, make([]int64, len(sp2) - len(sp1))...)
	} else if len(sp1) > len(sp2) {
		sp2 = append(sp2, make([]int64, len(sp1) - len(sp2))...)
	}

	i, l := 0, len(sp1)
	for i < l {
		if sp1[i] < sp2[i] {
			return -1
		}
		if sp1[i] > sp2[i] {
			return 1
		}
		i++
	}
	return 0
}