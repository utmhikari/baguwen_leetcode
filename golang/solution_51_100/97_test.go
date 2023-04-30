package solution_51_100

import "fmt"

//97. 交错字符串
//给定三个字符串s1、s2、s3，请你帮忙验证s3是否是由s1和s2 交错 组成的。
//
//两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：


var map97 map[string]bool

func solve97(s1 string, s2 string, s3 string, i1 int, i2 int) bool {
	l1, l2 := len(s1), len(s2)
	if i1 > l1 && i2 > l2{
		return false
	}

	key := fmt.Sprintf("%d_%d", i1, i2)
	_, ok := map97[key]
	if ok {
		return false
	}

	ans := false
	if i1 < l1 && s1[i1] == s3[i1 + i2] {
		if i2 == l2 && i1 + 1 == l1 {
			ans = true
		} else {
			ans = solve97(s1, s2, s3, i1 + 1, i2)
		}
	}
	if !ans{
		if i2 < l2 && s2[i2] == s3[i1 + i2] {
			if i1 == l1 && i2 + 1 == l2 {
				ans = true
			} else {
				ans = solve97(s1, s2, s3, i1, i2 + 1)
			}
		}
	}
	map97[key] = ans
	return ans
}


func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l1 + l2 != l3 {
		return false
	}
	if l1 == 0 {
		return s2 == s3
	}
	if l2 == 0 {
		return s1 == s3
	}
	map97 = make(map[string]bool)
	return solve97(s1, s2, s3, 0, 0)
}