package solution_51_100

import "fmt"

//87. 扰乱字符串
//使用下面描述的算法可以扰乱字符串 s 得到字符串 t ：
//如果字符串的长度为 1 ，算法停止
//如果字符串的长度 > 1 ，执行下述步骤：
//在一个随机下标处将字符串分割成两个非空的子字符串。
//即，如果已知字符串 s ，则可以将其分成两个子字符串 x 和 y ，且满足 s = x + y 。
//随机 决定是要「交换两个子字符串」还是要「保持这两个子字符串的顺序不变」。
//即，在执行这一步骤之后，s 可能是 s = x + y 或者 s = y + x 。
//在 x 和 y 这两个子字符串上继续从步骤 1 开始递归执行此算法。
//给你两个 长度相等 的字符串 s1 和s2，
//判断s2是否是s1的扰乱字符串。如果是，返回 true ；否则，返回 false


// false results
var cache87 map[string]bool

func isScrambleInternal(s1 string, s2 string, left1 int, right1 int, left2 int, right2 int) bool {
	if right1 - left1 != right2 - left2 {
		return false
	}
	key := fmt.Sprintf("%d_%d_%d_%d", left1, right1, left2, right2)
	fmt.Println(key)
	_, ok := cache87[key]
	if ok {
		return false
	}
	if left1 == right1 {
		if s1[left1] != s2[left2] {
			cache87[key] = true
			return false
		} else {
			return true
		}
	} else {
		i, j := left1, left2
		for i < right1 && j < right2 {
			if isScrambleInternal(s1, s2, left1, i, left2, j) &&
				isScrambleInternal(s1, s2, i + 1, right1, j + 1, right2) {
				return true
			}
			if isScrambleInternal(s1, s2, left1, i, right2 - (j - left2), right2) &&
				isScrambleInternal(s1, s2, i + 1, right1, left2, right2 - (j - left2) - 1) {
				return true
			}
			i++
			j++
		}
		cache87[key] = true
		return false
	}
}

func isScramble(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	if l1 != l2 {
		return false
	}
	cache87 = make(map[string]bool)
	ok := isScrambleInternal(s1, s2, 0, l1 - 1, 0, l2 - 1)
	fmt.Printf("%v\n", cache87)
	return ok
}